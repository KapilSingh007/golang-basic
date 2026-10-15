package main

import "fmt"

func main() {
	// Method 1 of defining variable (variable kind)

	var x int = 5
	fmt.Println("The value of x variable is", x)
	x = 15
	fmt.Println("The value of x variable after updating value is", x)

	// Method 2 of defining variable (constant kind)

	const y int = 5
	fmt.Println("The value of constant variable y is", y)

	// data type are
	// sign integer: int8, int16, int32, int64 or int
	// unsigned integer: uint8, uint16, uint32, uint64 or uint
	// string
	// array: [size]type

	// print statement
	// Println add /n or new line at the end of print statement
	fmt.Println(x)
	fmt.Println(y)

	// Printf used for formating string
	fmt.Printf("dfsafa sdfbsjh %v kjsbcfbsdjfb\n", x)

	// Declaring varibale

	// method 1: explicit way of declaring variable means we have to explicit about the data type
	var value1 int = 10

	fmt.Println(value1)

	// method 2: implicit way of declaring variable data type assign to variable according to value we are assigning

	value2 := 5

	fmt.Println(value2)
}
