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
