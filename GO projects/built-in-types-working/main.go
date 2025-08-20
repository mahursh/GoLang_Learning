package main

import (
	"fmt"
	"log"
)

// 1.basic types (numbers, strings, booleans)
var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64
var myUnit uint //only unsign integer .zero and positives
var myFloat float32
var myFloat64 float64
var myBool bool

// 2.aggerigate types (array, struct)
var myStringsArrays [3]string //In Go, you can only put declarations at the top level (outside main), not statements.

type MyStructCar struct {
	NumberOfTiers int
	Luxurt        bool
	BacketSeats   bool
	Make          string
	Model         string
	Year          int
}

// 3.reference types (pointers, slices, map, functions, channels)
// Pointer: pointer is nothing more that something that points to a specific location in memory.

// 4.interface type

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
	//------------------------------------------------------------
	myStringsArrays[0] = "cat"
	myStringsArrays[1] = "dog"
	myStringsArrays[2] = "mouse"

	fmt.Println("first element in array is : ", myStringsArrays[0])

	var myStructCar MyStructCar
	myStructCar.NumberOfTiers = 4
	myStructCar.Luxurt = true
	myStructCar.BacketSeats = false
	myStructCar.Make = "volkswagon"
	myStructCar.Model = "Something"

	myStructCar2 := MyStructCar{
		NumberOfTiers: 4,
		Luxurt:        true,
		BacketSeats:   false,
		Make:          "volvo",
		Model:         "xc90",
		Year:          1991,
	}

	fmt.Printf("MY CAR 1 IS A %d %s %s. \n", myStructCar.Year, myStructCar.Make, myStructCar.Model)
	fmt.Printf("MY CAR 2 IS A %d %s %s.\n", myStructCar2.Year, myStructCar2.Make, myStructCar2.Model)

	//------------------------------------------------------------

	x := 10
	myFirstPointer := &x
	fmt.Println()
	fmt.Println("X is :", x)
	fmt.Println("My first pointer is :", myFirstPointer)

	*myFirstPointer =  15
	fmt.Println("x is now", x)

}
