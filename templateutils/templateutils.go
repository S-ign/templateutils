package templateutils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/S-ign/stringutils"
)

// TemplateData names and organizes templates. Saving the template string to
// file in template/<catagory>/<name> format. Placeholder field used to easily
// indacate what fields to replace, and apply those changes to the template.
type TemplateData struct {
	Name         string
	Catagory     string
	Template     string
	PlaceHolders map[string]string
}

// String Pretty Print TemplateData Struct.
func (t *TemplateData) String() string {
	return fmt.Sprintf("Name: %v\nCatagory: %v\nTemplate: %v\nPlaceholders: %v",
		t.Name, t.Catagory, t.Template, t.PlaceHolders)
}

// ApplyPlaceholders updating t.PlaceHolders values before calling this method
// will replace all Placeholder values in the template.
// Ex. _________________________________________
// 								 <catagory> | <template_name> |
// t := GetTemplateData("test", "myname")       |
// t.PlaceHolders["name"] = John Doe            |
// t.ApplyPlaceholders()                        |
// template = Hello, my name is John Doe.       |
func (t *TemplateData) ApplyPlaceholders() string {
	tmpl, err := template.New(t.Name).Parse(t.Template)
	if err != nil {
		return ""
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, t.PlaceHolders)
	return output.String()
}

// UpdateTemplate updates the template text in its saved file location
// template/<catagory>/<name>
func (t *TemplateData) UpdateTemplate(template string) error {
	t.Template = template
	err := os.WriteFile(fmt.Sprintf("templates/%s/%s", t.Catagory, t.Name), []byte(template), 0644)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTemplate removes template text from saved location
// template/<catagory>/<name>
func (t *TemplateData) DeleteTemplate() {
	os.Remove(fmt.Sprintf("templates/%v/%v", t.Catagory, t.Name))
}

// ListTemplateCategories list directories in templates folder, which is used
// for the catagory names of the templates
func ListTemplateCategories() map[string][]string {
	files, err := ioutil.ReadDir("templates")
	if err != nil {
		return nil
	}

	catagories := make(map[string][]string)

	for _, f := range files {
		catagories["Catagories"] = append(catagories["Catagories"], f.Name())
	}

	return catagories
}

// ListTemplatesInCatagory Lists all templates within this catagory.
func ListTemplatesInCatagory(catagory string) (templates []*TemplateData, err error) {
	files, err := ioutil.ReadDir(fmt.Sprintf("templates/%v", catagory))
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		template, err := GetTemplateData(catagory, f.Name())
		if err != nil {
			return nil, err
		}
		templates = append(templates, template)
	}
	return templates, nil
}

// AddTemplateData Creates a new TemplateData struct, writing the template
// string to file, organizing it by template/<catagory>/<name>
// Template strings must have placeholders surrounded like {{.placeholder}}
// goodTemplateString := "Hello, my name is {{.name}}.
// badTemplateString := "Hello, my name is name.
func AddTemplateData(catagory, name, template string) (*TemplateData, error) {
	if !(strings.Contains(template, "{{.") && strings.Contains(template, "}}")) {
		return nil, fmt.Errorf(`error: template strings must contain placeholders
		surrounded by double curly braces
		ex. Template with {{.placeholder}}`)
	}

	path := fmt.Sprintf("templates/%v", catagory)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	path = fmt.Sprintf("templates/%v/%v", catagory, name)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.WriteFile(fmt.Sprintf("templates/%s/%s", catagory, name), []byte(template), 0644)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("error: template already exists, try updating or deleting this template instead")
	}
	placeholders := stringutils.GetAllSubstrInBetweenTwoStrs(template, "{{.", "}}")

	return &TemplateData{
		Name:         name,
		Catagory:     catagory,
		Template:     template,
		PlaceHolders: placeholders,
	}, nil
}

// GetTemplateData Retrives a premade template that was saved to file and
// returns a TemplateData so that its methods can be used.
func GetTemplateData(catagory, name string) (*TemplateData, error) {
	file, err := ioutil.ReadFile(fmt.Sprintf("templates/%v/%v", catagory, name))
	if err != nil {
		return nil, err
	}

	f := string(file)

	return &TemplateData{
		Name:         name,
		Catagory:     catagory,
		Template:     f,
		PlaceHolders: stringutils.GetAllSubstrInBetweenTwoStrs(f, "{{.", "}}"),
	}, nil
}
