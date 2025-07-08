package main

import (
	"context"
	"crud_postgres/controllers"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to DB
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Could not ping DB:", err)
	}
	fmt.Println("✅ Connected to PostgreSQL")

	// Set up router and controller
	r := chi.NewRouter()
	userController := &controllers.UserController{DB: db}

	r.Get("/user/{id}", userController.GetUser)
	r.Post("/user", userController.CreateUser)
	r.Put("/user/{id}", userController.UpdateUser)
	r.Delete("/user/{id}", userController.DeleteUser)

	// Run server
	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
