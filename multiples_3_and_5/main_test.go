package main

import "testing"

func TestMultiplesOfThree(t *testing.T) {
	result := multiplesOfThree([]int{1,2,3,4,5,6,7,8,9})

	if result != []int{3,6,9} {
		t.Error("Expected slice is %v", result)
	}
}
