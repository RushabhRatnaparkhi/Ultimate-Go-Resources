package main

import (
	"crypto/hmac"   // Provides functions for HMAC (keyed-hash message authentication code)
	"crypto/sha256" // Provides the SHA-256 hash algorithm
	"fmt"           // For formatted output
	"io"            // For writing string data to the hash function
)

func main() {
	// Generate and print HMAC for the first email
	c := getCode("test@example.com")
	fmt.Println(c)

	// Generate and print HMAC for a slightly different email
	c = getCode("test@exampl.com")
	fmt.Println(c)
}

// getCode takes a string input and returns its HMAC-SHA256 hash
func getCode(s string) string {
	// Create a new HMAC using SHA-256 and a secret key ("ourkey")
	h := hmac.New(sha256.New, []byte("ourkey"))

	// Write the input string into the HMAC
	io.WriteString(h, s)

	// Return the hexadecimal representation of the hash
	return fmt.Sprintf("%x", h.Sum(nil))
}
