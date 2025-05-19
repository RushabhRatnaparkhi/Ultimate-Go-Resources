package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// animal represents an animal with a name and a sound it makes
type animal struct {
	Name  string
	Sound string
}

// habitat represents the type of place where animals live
type habitat struct {
	Location string
	Type     string
	Climate  string
}

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml")) // Parses the template file
}

func main() {
	// Create individual animal instances
	a1 := animal{
		Name:  "Dog",
		Sound: "Bark",
	}

	a2 := animal{
		Name:  "Cat",
		Sound: "Meow",
	}

	a3 := animal{
		Name:  "Cow",
		Sound: "Moo",
	}

	// Create habitat instances
	h1 := habitat{
		Location: "Africa",
		Type:     "Savanna",
		Climate:  "Dry and Warm",
	}

	h2 := habitat{
		Location: "Amazon Rainforest",
		Type:     "Jungle",
		Climate:  "Humid and Tropical",
	}

	// Group animals and habitats into slices
	animals := []animal{a1, a2, a3}
	habitats := []habitat{h1, h2}

	// Create a composite data structure to pass into the template
	data := struct {
		Animals  []animal
		Habitats []habitat
	}{
		Animals:  animals,
		Habitats: habitats,
	}

	// Execute the template with the structured data
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err) // Exit if an error occurs
	}
}
