package main

import "scopes/packageOne"

var myVar = "myVar"

func main() {

	var blockVar = "blockVar"
	packageOne.PrintMe(blockVar)
	packageOne.PrintMe(myVar)
	packageOne.PrintMe(packageOne.PackageVar)

}
