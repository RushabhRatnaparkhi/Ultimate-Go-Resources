package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Start listening for TCP connections on port 8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err) // Fatal if the listener can't be started
	}
	defer li.Close() // Ensure the listener is closed when the program exits

	for {
		// Accept a new client connection
		conn, err := li.Accept()
		if err != nil {
			log.Println(err) // Log and continue on error
			continue
		}
		// Handle each connection in a new goroutine for concurrency
		go handle(conn)
	}
}

// handle processes the input from a single TCP client
func handle(conn net.Conn) {
	// Create a scanner to read input line-by-line from the connection
	scanner := bufio.NewScanner(conn)

	// Loop to read lines sent by the client
	for scanner.Scan() {
		ln := scanner.Text()   // Get the line as a string
		fmt.Println(ln)        // Print it to the server console
	}

	defer conn.Close() // Close the connection when done

	// The code below won’t be reached normally unless the client closes the connection
	// because the scanner reads until EOF or error.
	fmt.Println("Code got here.") // Indicates the connection has closed
}
