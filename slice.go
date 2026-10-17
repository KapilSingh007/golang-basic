package main

import "fmt"

func main() {
	var arr [10]int

	for i := 0; i < 10; i++ {
		arr[i] = i + 1
	}

	//slices

	// Method1: defining slice using array : array[startindex: uptoindex(not include)]
	s1 := arr[5:9]

	fmt.Printf("%T\n", s1)
	fmt.Println(s1, len(s1), cap(s1))

	// method2: by not creating array

	s2 := []int{1, 2, 3, 4, 5}

	// now append

	for i := 0; i < 2; i++ {
		s2 = append(s2, i+6)
	}

	fmt.Println(s2)

	//method3: make function

	s3 := make([]int, 10) // make(type, size,capacity)
	fmt.Println(s3)

	// NOTE: if array passes in a function it will passed as reference (means any change in slice value in function mutant original value )
}
