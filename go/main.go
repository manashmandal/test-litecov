package main

import "fmt"

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Subtract(5, 3))
	fmt.Println(Multiply(2, 3))
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// Uncovered function - no tests
func Power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// Another uncovered function
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}
