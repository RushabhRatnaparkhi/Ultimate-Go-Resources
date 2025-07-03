package controllers

import (
	"context"
	"crud_mongodb/models" // User model definition
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"        // Routing library
	"go.mongodb.org/mongo-driver/bson"           // MongoDB query language support
	"go.mongodb.org/mongo-driver/bson/primitive" // For working with MongoDB ObjectIDs
	"go.mongodb.org/mongo-driver/mongo"          // MongoDB driver
)

// UserController holds a MongoDB client instance
type UserController struct {
	client *mongo.Client
}

// Constructor: Returns a new UserController with the passed MongoDB client
func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client}
}

// GetUser handles GET /user/:id
// It finds a user by ObjectID and returns it as JSON
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") // Extract ":id" from URL

	// Convert string ID to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID format")
		return
	}

	// Set context with timeout to avoid hanging DB calls
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := uc.client.Database("go-web-dev-db").Collection("users")

	var user models.User

	// Find a user document by _id and decode into `user`
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "User not found")
		return
	}

	// Marshal user struct into JSON
	uj, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error marshalling user")
		return
	}

	// Write JSON to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uj)
}

// CreateUser handles POST /user
// It decodes JSON from request body and inserts a new user in MongoDB
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User

	// Decode request body JSON into user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid user data")
		return
	}

	// Generate a new MongoDB ObjectID for this user
	user.Id = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := uc.client.Database("go-web-dev-db").Collection("users")

	// Insert the user into the collection
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to insert user")
		return
	}

	// Marshal inserted user to JSON
	uj, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error marshalling user")
		return
	}

	// Send response with 201 Created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(uj)
}

// DeleteUser handles DELETE /user/:id
// It deletes a user from MongoDB based on their ID
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid ID format")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := uc.client.Database("go-web-dev-db").Collection("users")

	// Delete the user with the matching ObjectID
	res, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil || res.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "User not found or already deleted")
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted user with ID %s\n", id)
}
