package main

import "fmt"

func main() {

	// for loop
	n := 10
	for i := 1; i <= n; i++ {
		fmt.Printf("%v ", i)
	}

	//while loop
	i := 1
	for i <= n {
		fmt.Printf("%v ", i+10)
		i++
	}
}
