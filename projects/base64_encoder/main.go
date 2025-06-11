package main

import (
	"encoding/base64" // Package for base64 encoding and decoding
	"fmt"             // For printing to console
	"log"             // For logging fatal errors
)

func main() {
	// Define a plain-text message
	s := "Love is but a song to sing Fear's the way we die You can make the mountains ring Or make the angels cry Though the bird is on the wing And you may not know why Come on people now Smile on your brother Everybody get together Try to love one another Right now"

	// Encode the string into base64 format
	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64) // Print the base64-encoded string

	// Decode the base64 string back to original text
	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		// Handle decoding error and exit if any
		log.Fatalln("I'm giving her all she's got Captain!", err)
	}

	// Print the decoded string
	fmt.Println(string(bs))
}
