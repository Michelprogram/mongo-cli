package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var name string

type Entity struct {
	Name       string
	Attributes []*Attributes
}

func NewEntity(name string) *Entity {
	return &Entity{
		Name:       name,
		Attributes: make([]*Attributes, 0),
	}
}

func (e *Entity) PushAttributes(attributes *Attributes) {
	e.Attributes = append(e.Attributes, attributes)
}

type Attributes struct {
	Name string
	Type string
}

func NewAttributes(name, Type string) *Attributes {
	attribute := &Attributes{
		Name: "",
		Type: "",
	}

	attribute.SetName(name)
	attribute.SetType(Type)

	return attribute

}

func (a *Attributes) SetName(name string) {

	name = strings.ToLower(name)

	nameSplited := strings.Split(name, " ")

	for index, value := range nameSplited {
		nameSplited[index] = string(byte(value[0])-32) + value[1:]
	}

	a.Name = strings.Join(nameSplited, "")
}

func (a *Attributes) SetType(Type string) {
	a.Type = strings.ToLower(Type)
}

func (a Attributes) ToJsonFormat() string {

	var (
		res   = []byte(a.Name)
		blocs []string
		next  int
	)

	re := regexp.MustCompile("[A-Z]+")

	indexSplited := re.FindAllIndex(res, -1)

	for i := 0; i < len(indexSplited); i++ {

		index := indexSplited[i][0]

		if i+1 == len(indexSplited) {
			next = len(res)
		} else {
			next = indexSplited[i+1][0]
		}

		res[index] += 32

		blocs = append(blocs, string(res[index:next]))

	}

	return strings.Join(blocs, "_")
}

//Flags
func init() {

	flag.StringVar(&name, "name", "", "Name of your entity")

	flag.Parse()

	if name == "" {
		fmt.Println("You should specifie -name")
		os.Exit(1)
	}

}

//Ask questions
func AskAttributes() *Entity {

	var (
		attributes            *Attributes
		inputName, inputValue string
	)

	entity := *NewEntity(name)

	fmt.Printf("Start creating mongo entity : %s\n", name)
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

//Play with Template
func (e Entity) GenerateTemplate() error {

	name := "./generates/entity_" + e.Name + ".go"
	templateName := "./templates/entity.tmpl"

	file, err := os.Create(name)
	if err != nil {
		return err
	}

	template, err := template.ParseFiles(templateName)

	if err != nil {
		return err
	}

	template.Execute(file, e)

	return nil
}

//Generate files

func main() {

	//e := AskAttributes()

	e := NewEntity("Person")

	e.PushAttributes(NewAttributes("age", "int"))
	e.PushAttributes(NewAttributes("numero de telephone", "string"))

	e.GenerateTemplate()
}
