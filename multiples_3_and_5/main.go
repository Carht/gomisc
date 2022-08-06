package main

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

func sumMultiplesThreeAndFive(s []int) int {
	mulThree := multiplesOfThree(s)
	mulFive := multiplesOfFive(s)
	
	total := 0

	for _, i := range mulThree {
		total += i
	}

	for _, i := range mulFive {
		total += i
	}

	return total
}
