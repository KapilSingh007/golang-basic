package main

import "fmt"

type Sport struct {
	name string
}
type Person struct {
	name  string
	age   int
	sport Sport
}

func main() {
	var p1 Person = Person{name: "kapil", age: 24, sport: Sport{name: "football"}}

	fmt.Println(p1)
}
