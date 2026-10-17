package main

import "fmt"

func main() {

	// method 1: first intialize using var syntax then define value
	var arr1 [2]int
	arr1[0] = 5
	fmt.Println(arr1)

	// method2: implicite way

	arr2 := [2]int{1, 2}
	fmt.Println(arr2)

	// method2: implicite way but do not pass size
	arr3 := [...]int{3, 4}
	fmt.Println(arr3)

	// nested array

	arr4 := [3][2]int{{5, 6}, {7, 8}, {9, 10}}
	fmt.Println(arr4)

	// traversal of array
	// method1: formal loop traversal
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("%v ", arr2[i])
	}

	fmt.Println()

	//method2:

	for i, value := range arr3 {
		fmt.Printf("value at index %v is %v \n", i, value)
	}

	// NOTE: if array passes in a function it will passed as value (means any change in array value in function not mutant original value )
}
