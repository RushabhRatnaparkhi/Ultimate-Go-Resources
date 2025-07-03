package main

import (
	"context"
	"crud_mongodb/controllers"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := httprouter.New()

	client := getMongoClient()
	userController := controllers.NewUserController(client)

	// ✅ Wrap GetUser with ObjectID validation middleware
	r.GET("/user/:id", userController.GetUser)

	r.POST("/user", userController.CreateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}

func getMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB Ping Failed:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")
	return client
}
