package main

import (
	"github.com/Michelprogram/mongo-cli/mongogo"
)

func main() {

	//e := AskAttributes()

	e := NewEntity("Person")

	e.PushAttributes(NewAttributes("age", "int"))
	e.PushAttributes(NewAttributes("numero de telephone", "string"))
	e.PushAttributes(NewAttributes("BeeKeeper", "Beekeeper"))

	e.GenerateTemplate()
}
