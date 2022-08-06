package main

import "fmt"

func multiplesOfThree(s []int) []int {
	mulThree := []int{}

	for _, i := range s {
		if i % 3 == 0 {
			mulThree = append(mulThree, i)
		}
	}

	return mulThree
}

func multiplesOfFive(s []int) []int {
	mulFive := []int{}

	for _, i := range s {
		if i % 5 == 0 {
			mulFive = append(mulFive, i)
		}
	}

	return mulFive
}

func sumMultiplesThreeAndFive(n int) int {
	s := make([]int, n + 1)
	total := 0

	for i := 0; i < n; i++ {
		s[i] = i
	}

	multThree := multiplesOfThree(s)
	multFive := multiplesOfFive(s)

	for _, i := range multThree {
		total += i
	}

	for _, i := range multFive {
		total += i
	}

	return total
}

func main() {
	sumBelow1k := sumMultiplesThreeAndFive(1000)
	fmt.Printf("The sum for the multiples of three and five is: %d\n", sumBelow1k)
}
