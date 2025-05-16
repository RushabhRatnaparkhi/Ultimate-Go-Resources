package main

import (
	"log"           // For logging fatal errors
	"os"            // For writing output to the terminal
	"text/template" // For parsing and executing HTML/text templates
)

// tpl holds all parsed templates from the "templates" directory
var tpl *template.Template

// init is automatically run before main()
// Used here to parse all .gohtml templates at program startup
func init() {
	// Parse all template files inside the "templates" folder
	// template.Must wraps the parsing and panics if there's an error
	// This ensures the app fails fast if templates are broken
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// Execute the template named "template.gohtml"
	// Passing the string "awesome" as data into the template
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", "awesome")
	if err != nil {
		log.Fatalln(err) // Log the error and exit the program
	}
}
