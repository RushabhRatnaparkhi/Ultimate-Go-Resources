package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// Declare a global template variable
var tpl *template.Template

// Define a struct representing a sage with a name and motto
type sage struct {
	Name  string
	Motto string
}

// Define a struct representing a car with manufacturer, model, and door count
type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// Create a FuncMap to register custom functions to be used inside the template.
// "uc" will call strings.ToUpper to convert a string to uppercase.
// "ft" will use a custom function to return the first three characters of a string.
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	// Parse the template file and register the FuncMap.
	// Why ? Because custom functions must be defined before parsing the template file that uses them.
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template.gohtml"))
}

// firstThree returns the first 3 characters of a trimmed string
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	// Create instances of sages
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	// Create instances of cars
	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}
	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	// Combine sages and cars into slices
	sages := []sage{b, g, m}
	cars := []car{f, c}

	// Group data into a single struct to pass to the template
	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sages,
		Transport: cars,
	}

	// Execute the template with the provided data
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
