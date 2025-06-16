package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define a struct to represent the thumbnail details
type thumbnail struct {
	URL           string // URL of the thumbnail image
	Height, Width int    // Dimensions of the thumbnail
}

// Define a struct to represent the main image and its metadata
type img struct {
	Width, Height int       // Dimensions of the main image
	Title         string    // Title of the image
	Thumbnail     thumbnail // Embedded thumbnail struct
	Animated      bool      // Flag to indicate if the image is animated
	IDs           []int     // A list of associated IDs
}

func main() {
	// Declare a variable of type img to hold the parsed JSON data
	var data img

	// JSON string received (as if from an API or file)
	rcvd := `{
		"Width":800,
		"Height":600,
		"Title":"View from 15th Floor",
		"Thumbnail":{
			"Url":"http://www.example.com/image/12345",
			"Height":125,
			"Width":100
		},
		"Animated":false,
		"IDs":[116,943,234,38793]
	}`

	// Use json.Unmarshal to parse the JSON string into the Go struct.
	// IMPORTANT: We pass &data (a pointer to the data variable), because
	// Unmarshal needs to modify the actual memory location of 'data'.
	// If we just passed 'data' (a copy), changes would not be saved.
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}

	// Print the full struct to see all the parsed data
	fmt.Println(data)

	// Loop through the IDs and print each with its index
	for i, v := range data.IDs {
		fmt.Println(i, v)
	}

	// Print the URL of the thumbnail image
	fmt.Println(data.Thumbnail.URL)
}
