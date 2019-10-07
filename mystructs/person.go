package mystructs

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

func (p Person) Speak() {
	fmt.Printf("I am %s, %s. and Im %d years old\n", p.First, p.Last, p.Age)
}
