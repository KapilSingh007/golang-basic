package main

import "fmt"

func main() {

	// method1: using var

	var mp1 map[string]int = map[string]int{}

	mp1["b"] = 2
	fmt.Println(mp1)

	// method2
	mp2 := map[string]int{}

	mp2["b"] = 2
	fmt.Println(mp2)

	// if key exist or not

	value, ok := mp2["b"]

	fmt.Println(value, ok)
}
