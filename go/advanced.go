package main

// This entire file has NO tests - should be flagged by litecov

// GCD calculates the greatest common divisor
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM calculates the least common multiple
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fibonacci returns the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
