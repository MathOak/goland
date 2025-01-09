package main

import (
	"testing"
)

func ShouldAddTwoNumbers(t *testing.T) {
	if add(1, 2) != 3 {
		t.Error("Expected 1 + 2 to equal 3")
	}
	if add(-1, -1) != -2 {
		t.Error("Expected -1 + -1 to equal -2")
	}
}

func ShouldSubtractTwoNumbers(t *testing.T) {
	if sub(2, 1) != 1 {
		t.Error("Expected 2 - 1 to equal 1")
	}
	if sub(-1, -1) != 0 {
		t.Error("Expected -1 - -1 to equal 0")
	}
}

func ShouldMultiplyTwoNumbers(t *testing.T) {
	if mult(2, 3) != 6 {
		t.Error("Expected 2 * 3 to equal 6")
	}
	if mult(-1, -1) != 1 {
		t.Error("Expected -1 * -1 to equal 1")
	}
}

func ShouldDivideTwoNumbers(t *testing.T) {
	result, err := div(6, 3)
	if err != nil {
		t.Error("Expected no error")
	}
	if result != 2 {
		t.Error("Expected 6 / 3 to equal 2")
	}

	_, err = div(1, 0)
	if err == nil {
		t.Error("Expected error when dividing by zero")
	}
}
