package main

import (
	"log"

	// helpers "packages/helpers"
	helpers "github.com/devj1003/mypackage/helpers"
)

func main() {
	log.Println("Hello!")

	var MyVar helpers.SomeType
	MyVar.TypeName = "Some Name"
	log.Println(MyVar.TypeName)
}
