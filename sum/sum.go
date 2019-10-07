// Package sum provides ...
package sum

func SumofEven(s func(i ...int) int, j ...int) int {
	var even []int
	for _, x := range j {
		if IsEven(x) {
			even = append(even, x)
		}
	}
	return s(even...)
}

func IsEven(i int) bool {
	return (i%2 == 0)
}

func GetOddList(i ...int) []int {
	var odd []int
	for _, v := range i {
		if !IsEven(v) {
			odd = append(odd, v)
		}
	}
	return odd
}

func GetEvenList(i ...int) []int {
	var even []int
	for _, v := range i {
		if IsEven(v) {
			even = append(even, v)
		}
	}
	return even
}

// Func as param.
// Example of func as param.
func Sumof(sl func(...int) []int, l ...int) int {
	return Sum(sl(l...)...)
}

// Sum accepting variatic parameter,
// which becomes slice of int
func Sum(i ...int) int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}
