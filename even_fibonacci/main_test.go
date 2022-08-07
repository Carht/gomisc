package main

import "testing"

func TestFib(t *testing.T) {
	result := fib(7)
	expected := 13

	if result != expected {
		t.Errorf("The expected value is %d, but the result was %d", expected, result)
	}
}

func TestFib27(t *testing.T) {
	result := fib(27)
	expected := 196418

	if result != expected {
		t.Errorf("The expected value is %d, but the result was %d", expected, result)
	}
}
