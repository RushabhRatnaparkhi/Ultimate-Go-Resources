package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// Start a TCP listener on port 8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err) // Terminate the program if unable to bind to the port
	}
	defer li.Close() // Ensure the listener is closed when the main function exits

	for {
		// Accept incoming client connections
		conn, err := li.Accept()
		if err != nil {
			log.Println(err) // Log any connection errors and continue accepting new connections
			continue
		}
		// Handle the connection in a new goroutine (concurrent client handling)
		go handle(conn)
	}
}

// handle processes the incoming data from a single TCP client
func handle(conn net.Conn) {
	// Use a scanner to read data line-by-line from the connection
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		// Read the input line, convert it to lowercase
		ln := strings.ToLower(scanner.Text())

		// Convert the line to a byte slice
		bs := []byte(ln)

		// Apply ROT13 transformation to the byte slice
		r := rot13(bs)

		// Send both the original and transformed strings back to the client
		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
	}
	// Connection will be closed automatically by the caller once this goroutine ends
}

// rot13 performs a basic ROT13 cipher transformation on the input byte slice
func rot13(bs []byte) []byte {
	// Allocate a new byte slice for the transformed result
	var r13 = make([]byte, len(bs))

	for i, v := range bs {
		// Apply ROT13 logic for lowercase ASCII letters (a-z → ASCII 97–122)
		if v <= 109 {
			// 'a' to 'm' shift forward by 13
			r13[i] = v + 13
		} else {
			// 'n' to 'z' shift backward by 13
			r13[i] = v - 13
		}
	}

	return r13
}
