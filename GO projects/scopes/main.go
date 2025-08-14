package main

import "scopes/packageOne"

var myVar = "myVar"

func main() {

	var blockVar = "blockVar"
	packageOne.PrintMe(blockVar)
	packageOne.PrintMe(myVar)
	packageOne.PrintMe(packageOne.PackageVar)

	//in the main function, print out the values of myVar,
	//blockVar, and PackageVar on one line using the PrintMe  function in packageOne

}
