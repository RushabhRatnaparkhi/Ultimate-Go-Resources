package main

import (
	"log"           // For logging errors
	"os"            // For accessing system resources like files and stdout
	"text/template" // For working with plain text templates (as opposed to html/template for web)
)

func main() {
	// Load and parse one or more template files.
	// You can pass multiple filenames here: template.ParseFiles("file1.gohtml", "file2.gohtml", ...)
	// Conventionally, Go uses `.gohtml` as the file extension for template files to distinguish them from regular HTML.
	template, err := template.ParseFiles("template.gohtml")
	if err != nil {
		// log.Fatalln prints the error and exits the program
		log.Fatalln(err)
	}

	// Execute the parsed template and write the output to os.Stdout (i.e., console)
	// The second argument is the data to be injected into the template; `nil` means no data is passed
	err = template.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// DONT USE THE ABOVE CODE IN PRODUCTION, MORE IMPROVEMENTS TO BE MADE

