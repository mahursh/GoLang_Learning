package main

import "fmt"

func main() {
	//var   name   type
	//var whatToSay string
    //whatToSay = "Hello, World!"


	// := is the short variable declaration operator that defines and initializes a variable in one step.
	whatToSay := "Hello, World!"
	sayHelloWorld(whatToSay)
}

func sayHelloWorld(whatToSay string){
	fmt.Println(whatToSay)
}
