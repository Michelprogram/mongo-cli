package mongogo

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

var name string

type Entity struct {
	Name       string
	Attributes []*Attributes
}

func NewEntity(name string) *Entity {

	name = strings.ToUpper(string(name[0])) + name[1:]

	return &Entity{
		Name:       name,
		Attributes: make([]*Attributes, 0),
	}
}

func (e *Entity) PushAttributes(attributes *Attributes) {
	e.Attributes = append(e.Attributes, attributes)
}

//Play with Template
func (e Entity) GenerateTemplate() error {

	path, err := os.Getwd()

	if err != nil {
		return err
	}

	name := path + "/entity_" + e.Name + ".go"
	templateName := "template_entity.tmpl"

	file, err := os.Create(name)
	if err != nil {
		return err
	}

	template, err := template.New(templateName).Parse(GetTemplate())

	if err != nil {
		return err
	}

	template.Execute(file, e)

	return nil
}

//Ask questions
func AskAttributes() *Entity {

	var (
		attributes            *Attributes
		inputName, inputValue string
	)

	entity := *NewEntity(name)

	fmt.Printf("Start creating mongo entity : %s\n", name)
	fmt.Println("Basic type are int, string, bool")
	fmt.Println("When you have finish press type : done")

	for {

		fmt.Printf("Attributes name : ")
		fmt.Scan(&inputName)

		if inputName == "done" {
			break
		}

		fmt.Printf("Attributes type : ")
		fmt.Scan(&inputValue)

		if inputName == "done" {
			break
		}

		attributes = NewAttributes(inputName, inputValue)

		entity.PushAttributes(attributes)
	}

	return &entity

}

//Flags
func init() {

	flag.StringVar(&name, "name", "", "Name of your entity")

	flag.Parse()

	if name == "" {
		fmt.Println("You should specifie flag -name.")
		os.Exit(1)
	}

}
