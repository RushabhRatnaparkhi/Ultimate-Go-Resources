package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Define a person struct to be converted to JSON
type Person struct {
	FirstName string   `json:"fname"`
	LastName  string   `json:"lname"`
	Items     []string `json:"items"`
}

func main() {
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/marshal", handleMarshal)
	http.HandleFunc("/encode", handleEncode)

	// Start the server on port 8080
	log.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// handleHomePage serves a basic HTML response
func handleHomePage(w http.ResponseWriter, req *http.Request) {
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Home</title>
	</head>
	<body>
		<p>You are at the homepage</p>
	</body>
	</html>`
	w.Write([]byte(html))
}

// handleMarshal converts a Go struct to JSON using json.Marshal and writes it to the response
func handleMarshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	person := Person{
		FirstName: "James",
		LastName:  "Bond",
		Items:     []string{"Suit", "Gun", "Wry sense of humor"},
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return
	}

	w.Write(jsonData)
}

// handleEncode encodes a Go struct to JSON and directly writes it to the HTTP response using json.Encoder
func handleEncode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	person := Person{
		FirstName: "James",
		LastName:  "Bond",
		Items:     []string{"Suit", "Gun", "Wry sense of humor"},
	}

	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		log.Println("Error encoding JSON:", err)
	}
}