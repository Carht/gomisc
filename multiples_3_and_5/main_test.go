package main

import "testing"

func TestMultiplesOfThree(t *testing.T) {
	result := multiplesOfThree([]int{3, 6, 9})
	expected := []int{3, 6, 9}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected slice is %v", result)
		}
	}
}

func testMultiplesOfThree(t *testing.T) {
	result := multiplesOfThree([]int{1, 3, 6, 9})
	expected := []int{3, 6, 9}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected slice is %v", result)
		}
	}
}

func testMultipleOfThree(t *testing.T) {
	result := multiplesOfThree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	expected := []int{3, 6, 9, 12, 15, 18}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected slice is %v", result)
		}
	}
}

