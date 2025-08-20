package main

import "log"

// basic types (numbers, strings, booleans)
var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64
var myUnit uint //only unsign integer .zero and positives
var myFloat float32
var myFloat64 float64

// aggerigate types (array, struct)

// reference types (pointers, slices, map, functions, channels)

// interface type
func main() {

	myInt = 10
	myUnit = 20
	myFloat = 10.1
	myFloat = 100.1
	log.Println(myInt,myUnit,myFloat,myFloat64)
}
