package main

import "fmt"

func main() {

	// if else conditional statement
	x := 6

	if x > 5 {
		fmt.Println("variable x is greater than 5")
	} else {
		fmt.Println("variable x is not greater than 5")
	}

	// switch case conditional statement

	// one switch based on just passing value
	switch x {
	case 5:
		fmt.Println("variable x is equal to 5")
	default:
		fmt.Println("variable x is not equal than 5")
	}

	// second switch is

	switch {
	case x > 5:
		fmt.Println("variable x is greater than 5")
	case x == 5:
		fmt.Println("variable x is equal to 5")
		fallthrough // this is used for run code below this condtion and not break immediately
	default:
		fmt.Println("variable x is less than 5")
	}
}
