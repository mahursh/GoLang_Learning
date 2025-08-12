package main

import (
	"fmt"
	"scopes/packageOne"
)

var one = "One" // package level variable

func main() {
	// block level variable
	var one = "this is a block level variable" // variable shadowing
	fmt.Println(one)
	myFunc()

	newString := packageOne.PublicVar
	fmt.Println("from packageOne : ", newString)

	packageOne.Exported()

}

func myFunc() {

	fmt.Println(one)

}
