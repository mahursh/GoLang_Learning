package main

import (
	"fmt"
	"log"
)

// basic types (numbers, strings, booleans)
var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64
var myUnit uint //only unsign integer .zero and positives
var myFloat float32
var myFloat64 float64
var myBool bool

// aggerigate types (array, struct)
var myStringsArrays [3]string //In Go, you can only put declarations at the top level (outside main), not statements.





// reference types (pointers, slices, map, functions, channels)

// interface type
func main() {

	myInt = 10
	myUnit = 20
	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUnit, myFloat, myFloat64)

	myString := "one"
	log.Println(myString)
	myString = "two"
	log.Println(myString)
	//Strings in GO are imutable. that means that when you try to change the value of a string , what you are
	//actually doing is creating an entirely new String and storing that in your variable.

	myBool = true //the only thing you can strore in a boolean variable is true and false. we cant't even put 0 in.
	log.Println(myBool)

	
	myStringsArrays[0] = "cat"
	myStringsArrays[1] = "dog"
	myStringsArrays[2] = "mouse"

	fmt.Println("first element in array is : ", myStringsArrays[0])

}
