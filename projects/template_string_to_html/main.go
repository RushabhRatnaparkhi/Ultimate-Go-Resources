package main

import (
	"log"           // For logging fatal errors
	"os"            // For file and output handling
	"text/template" // For parsing and executing plain text templates
)

// Declare a global template pointer
var tpl *template.Template

// init() is automatically executed before main()
// Use it to parse all template files in the "templates" folder
func init() {
	// Parse all files in the "templates" directory (e.g., *.gohtml)
	// template.Must automatically handles errors: it panics if parsing fails
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// 👇 This executes the default template — which is the first parsed file alphabetically.
	// If "one.gohtml" is the first file in that order, it will be rendered here.
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err) // Print and exit on error
	}

	// Execute the specific template named "template.gohtml"
	err = tpl.ExecuteTemplate(os.Stdout, "template.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Execute the template named "one.gohtml"
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Execute the template named "two.gohtml"
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
