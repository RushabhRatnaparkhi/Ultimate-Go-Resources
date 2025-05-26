package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Start listening on TCP port 8080 on all available interfaces.
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		// If the server fails to start, log the error and exit the program.
		log.Fatalln(err)
	}
	// Ensure the listener is closed when the main function exits.
	defer li.Close()

	for {
		// Wait for and accept an incoming TCP connection.
		conn, err := li.Accept()
		if err != nil {
			// If there's an error accepting the connection, log it and continue to accept the next one.
			log.Println(err)
			continue
		}

		// Write a welcome message to the client using different output functions.
		io.WriteString(conn, "\nHello from TCP server\n") // Basic message using io
		fmt.Fprintln(conn, "How is your day?")            // Prints with a newline
		fmt.Fprintf(conn, "%v", "Well, I hope!")          // Prints formatted string

		// Close the connection after responding.
		conn.Close()
	}
}
