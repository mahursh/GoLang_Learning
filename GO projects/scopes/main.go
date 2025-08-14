package main

import "scopes/packageOne"

var myVar = "myVar"

func main() {

	var blockVar = "blockVar"
	packageOne.PrintMe(blockVar, myVar, packageOne.PackageVar)
	packageOne.PrintMe2(blockVar, myVar, packageOne.PackageVar)

}
