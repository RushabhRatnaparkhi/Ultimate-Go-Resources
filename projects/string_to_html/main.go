package main

import (
	"fmt"     // For formatted I/O
	"io"      // For copying data between sources (e.g., readers and writers)
	"log"     // For logging errors
	"os"      // For working with the file system and command-line arguments
	"strings" // For working with string readers
)

func main() {
	// Get the second command-line argument (index 1), which is expected to be the user's name
	name := os.Args[1]

	// Print the name of the program (index 0)
	fmt.Println(os.Args[0])

	// Print the name argument passed in (index 1)
	fmt.Println(name)

	// Create a string containing HTML content, including the user's name
	// Using fmt.Sprint to safely build a multiline string
	htmlContent := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Hello World!</title>
		</head>
		<body>
			<h1>` + name + `</h1>
		</body>
		</html>
	`)

	// Create (or overwrite) a file named "index.html"
	htmlFile, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file:", err)
	}
	// Ensure the file is closed properly after writing
	// defer is used to postpone the execution of a function until the surrounding function (main() in this case) returns.
	defer htmlFile.Close()

	// Write the HTML content to the file using io.Copy and a string reader
	io.Copy(htmlFile, strings.NewReader(htmlContent))
}