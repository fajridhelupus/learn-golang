package main

import "fmt"

func main() {
	//MAIN Types
	//string
	//bool
	//int
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//byte-alias for uint8
	//rune-alias for int32
	//float32 float64
	//complex63 complex128

	//using var
	// var name = "fajri"
	var age int32 = 16
	var isCool = true
	var size float32 = 0.5

	//Shorthand
	// name := "fajri"
	// email := "fajri@gmail.com"

	name, email := "fajri", "fajri@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)
}
