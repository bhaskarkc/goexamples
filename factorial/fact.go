package factorial

// Recursive function.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// Same function with loop.
func LoopFact(n int) int {
	t := 1
	for ; n > 0; n-- {
		t *= n
	}
	return t
}
