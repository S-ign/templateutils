package main

import (
	"fmt"

	"github.com/S-ign/templateutils/templateutils"
)

func main() {
	// add template
	temp, err := templateutils.AddTemplateData("test", "myname", "Hello everyone! My name is {{.name}}, nice to meet you!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(temp)

	// replace placeholders with data and apply it to the current template
	temp.PlaceHolders["name"] = "Martin"
	temp.PlaceHolders["day"] = "Monday"
	fmt.Println(temp.ApplyPlaceholders())

	fmt.Println("--------------------------------------------------")
	// get list of all template catagories
	fmt.Println(templateutils.ListTemplateCategories())
	fmt.Println("--------------------------------------------------")

	// list all templates in a catagory
	templates, err := templateutils.ListTemplatesInCatagory("test")
	if err != nil {
		fmt.Println(err)
	}
	for _, t := range templates {
		fmt.Println(t.Template)
	}
	fmt.Println("--------------------------------------------------")

	// update an existing template
	fmt.Println("Updating Template...")
	temp.UpdateTemplate("Whats up everyone! My name is {{.name}}, nice to meet you!")
	fmt.Println(temp)

	temp.DeleteTemplate()
	fmt.Println(temp)
	fmt.Println("--------------------------------------------------")
	fmt.Println("--------------------------------------------------")
	fmt.Println("--------------------------------------------------")
	templateutils.ListAllTemplates()
}
