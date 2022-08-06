package main

import "testing"

func TestMultiplesOfThreeDirectReturn(t *testing.T) {
	result := multiplesOfThree([]int{3, 6, 9})
	expected := []int{3, 6, 9}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("The expected is %v, but the answer was %v", expected, result)
		}
	}
}

func TestMultiplesOfThreeAditionalArgument(t *testing.T) {
	result := multiplesOfThree([]int{1, 3, 6, 9})
	expected := []int{3, 6, 9}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("The expected is %v, but the answer was %v", expected, result)
		}
	}
}

func TestMultipleOfThreeMultiplesArguments(t *testing.T) {
	result := multiplesOfThree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	expected := []int{3, 6, 9, 12, 15, 18}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("The expected is %v, but the answer was %v", expected, result)
		}
	}
}

func TestMultiplesOfFive(t *testing.T) {
	result := multiplesOfFive([]int{5, 10})
	expected := []int{5, 10}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("The expected is %v, but the answer was %v", expected, result)
		}
	}
}

func TestMultiplesOfFiveAditionalArgument(t *testing.T) {
	result := multiplesOfFive([]int{1, 5, 10})
	expected := []int{5, 10}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("The expected is %v, but the answer was %v", expected, result)
		}
	}
}

func TestMultiplesOfFiveMultiplesArguments(t *testing.T) {
	result := multiplesOfFive([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	expected := []int{5, 10, 15, 20}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("The expected is %v, but the answer was %v", expected, result)
		}
	}
}

func TestSumNumbersSlice(t *testing.T) {
	result := sumMultiplesThreeAndFive(10)
	expected := 23

	if result != expected {
		t.Errorf("The expected is %d, but the answer was %v", expected, result)
	}
}
