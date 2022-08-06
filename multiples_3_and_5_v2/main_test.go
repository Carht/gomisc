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

func TestSumM(t *testing.T) {
	result := sumM(10)
	expected := 23

	if result != 23 {
		t.Errorf("The expected value is %d, but the answer was %d", expected, result)
	}
}

func TestSumGauss(t *testing.T) {
	result := sumM(1000)
	expected := 233168

	if result != expected {
		t.Errorf("The expected value is %d, but the answer was %d", expected, result)
	}
}
