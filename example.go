package main

import (
	"fmt"

	"github.com/Michelprogram/mongo-cli/mongogo"
)

func main() {

	e := mongogo.AskAttributes()

	err := e.GenerateTemplate()

	if err != nil {
		fmt.Println("Err : ", err)
	}
}
