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

func TestFibListMinorsOF4_000_000(t *testing.T) {
	result := fibMinor4m()
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144,
		233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657,
		46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269,
		2178309, 3524578}

	if len(result) != len(expected) {
		t.Errorf("The lengs of the slices are diferent, the expected leng is %d", len(expected))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("The expected list is %v, but the answer was %v", expected, result)
		}
	}
}

func TestGetEvenNumbers(t *testing.T) {
	result := evenp(4)
	expected := true

	if expected != true {
		t.Errorf("The expected value is %t, but the answer was %t", expected, result)
	}
}

func TestGetEvenNumbersWithNotEvenNumbers(t *testing.T) {
	result := evenp(7)
	expected := false

	if expected != false {
		t.Errorf("The expected value is %t, but the answer was %t", expected, result)
	}
}

func TestGetEvenElementsFromASlice(t *testing.T) {
	result := evenSlicep([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	expected := []int{2, 4, 6, 8, 10}

	if len(result) != len(expected) {
		t.Errorf("The len of expected is %d, but the len of result was %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("The expected slice is %v, but the answer was %v", expected, result)
		}
	}
}

func TestSumElementsOfASlice(t *testing.T) {
	result := sumS([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	expected := 55

	if result != expected {
		t.Errorf("The expected value is %d, but the answer was %d", expected, result)
	}
}
