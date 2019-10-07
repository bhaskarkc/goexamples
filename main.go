package main

import (
	"./factorial"
	"./mystructs"
	"./sum"
	"fmt"
)

func main() {

	mystructs.Methodset()

	mystructs.UnmarshallExample(mystructs.MarshallExample)
	return

	// Factorial
	fmt.Println(factorial.Factorial(3))
	fmt.Println(factorial.LoopFact(3))
	// when I no longer wanted to be different. It is different.
	// Be 100% focused and be 100% relaxed. #swimming.

	// Sum
	ii := []int{1, 2, 3, 4, 5}
	s := sum.Sum(ii...)
	fmt.Printf("Sum of each int items in a given list: %v\n", s)

	// Sum of odd.
	ol := sum.GetOddList(ii...)
	so := sum.Sum(ol...)
	fmt.Printf("Odd numbers %v\n", ol)
	fmt.Printf("Sum of Odd numbers: %v\n", so)

	// Sum of even.
	// Example of passing func as argument.
	se := sum.Sumof(sum.GetEvenList, ii...)
	fmt.Printf("Even numbers %v\n", sum.GetEvenList(ii...))
	fmt.Printf("Sum of Even numbers: %v\n", se)
}
