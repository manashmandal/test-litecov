package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Add(1, 2) should be 3")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(5, 3) != 2 {
		t.Error("Subtract(5, 3) should be 2")
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(2, 3) != 6 {
		t.Error("Multiply(2, 3) should be 6")
	}
}

// Note: Divide, Power, Factorial are NOT tested - should show as uncovered
