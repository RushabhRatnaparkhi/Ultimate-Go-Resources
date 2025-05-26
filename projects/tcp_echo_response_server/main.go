package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Start listening on TCP port 8080 on all interfaces.
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err) // Exit if listener setup fails
	}
	defer li.Close() // Ensure listener is closed when the program ends

	for {
		// Accept an incoming client connection
		conn, err := li.Accept()
		if err != nil {
			log.Println(err) // Log the error and continue accepting new connections
			continue
		}

		// Handle the connection in a new goroutine to support concurrency
		go handle(conn)
	}
}

// handle processes communication with a single client
func handle(conn net.Conn) {
	// Create a scanner to read input from the connection line by line
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()                         // Read the line of input
		fmt.Println(ln)                              // Print it on the server console
		fmt.Fprintf(conn, "I heard you say: %s\n", ln) // Echo the line back to the client
	}

	// Close the connection after client disconnects or stream ends
	defer conn.Close()

	// This line executes only if scanner exits (e.g., client disconnects)
	fmt.Println("Code got here.")
}
