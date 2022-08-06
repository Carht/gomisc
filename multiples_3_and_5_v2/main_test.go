package main

import "testing"

func TestMultiple(t *testing.T) {
	result := mult(3, 3)
	expected := true

	if result != expected {
		t.Errorf("The expected value is %t, but the answer was %t", expected, result)
	}
}

func TestMultipleForFive(t *testing.T) {
	result := mult(3, 5)
	expected := false

	if result != expected {
		t.Errorf("The expected value is %t, but the answer was %t", expected, result)
	}
}

func TestMultipleForTwo(t *testing.T) {
	result := mult(12, 2)
	expected := true

	if result != expected {
		t.Errorf("The expected value is %t, but the answer was %t", expected, result)
	}
}
