package mongogo

import (
	"regexp"
	"strings"
)

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

func (a Attributes) DisplayType() string {

	switch a.Type {
	case "int", "string", "bool":
		return a.Type
	default:
		return "*" + string(a.Type[0]-32) + a.Type[1:]
	}

}
