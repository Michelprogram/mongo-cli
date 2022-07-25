package mongogo

import (
	"flag"
	"fmt"
)

//Flags
func init() {
	var params, name string

	flag.StringVar(&params, "generate", "", "Generate mongoDB structure")
	flag.StringVar(&name, "name", "", "Name of your entity")

	flag.Parse()

}

//Play with Template

//Generate files
func Test() {
	fmt.Println("Test")
}
