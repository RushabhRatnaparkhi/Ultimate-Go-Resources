package main

import (
	"html/template" // For safe HTML rendering
	"log"           // For logging errors
	"os"            // For output to terminal
)

// Page struct defines the data to be passed to the HTML template
type Page struct {
	Title   string // The title of the page
	Heading string // A heading to display
	Input   string // User input that may contain unsafe content
}

var tpl *template.Template // Global template variable

// init function runs before main and parses the template file
func init() {
	// Parse the HTML template file named "template.gohtml"
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	// Data to render into the template
	home := Page{
		Title:   "Safe Web Page", // Changed title
		Heading: "Security with html/template", // Updated heading
		Input:   `<b>This should not be bold!</b><script>alert("Hacked!");</script>`, // Potentially unsafe input
	}

	// Execute the template and write the result to standard output
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", home)
	if err != nil {
		log.Fatalln(err) // Log and exit on error
	}
}