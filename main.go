package main

import (
	"fmt"

	templateutils "github.com/S-ign/templateutils/utils"
)

func main() {
	temp, err := templateutils.AddTemplateData("test", "myname", "Hello, my name is {{.name}}, {{.day}}'s are always delightful.")
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
}
