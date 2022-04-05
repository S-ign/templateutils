package main

import (
	"fmt"

	"github.com/S-ign/templateutils/templateutils"
)

func main() {
	temp, err := templateutils.AddTemplateData("test", "myname", "Hello everyone! My name is {{.name}}, nice to meet you!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(temp)

	temp.PlaceHolders["name"] = "Martin"
	temp.PlaceHolders["day"] = "Monday"
	fmt.Println(temp.ApplyPlaceholders())

	fmt.Println("--------------------------------------------------")
	fmt.Println(templateutils.ListTemplateCategories())
	fmt.Println("--------------------------------------------------")

	templates, err := templateutils.ListTemplatesInCatagory("test")
	if err != nil {
		fmt.Println(err)
	}
	for _, t := range templates {
		fmt.Println(t.Template)
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("Updating Template...")
	temp.UpdateTemplate("Whats up everyone! My name is {{.name}}, nice to meet you!")
	fmt.Println(temp)

	temp.DeleteTemplate()
	fmt.Println(temp)
}
