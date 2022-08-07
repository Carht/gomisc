package main

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	} else {
		return fib(n - 1) + fib(n - 2)
	}
}

func fibMinor4m() []int {
	minor4m := []int{}

	for i := 0; i < 100; i++ {
		fibLocal := fib(i)
		if fibLocal < 4000000 {
			minor4m = append(minor4m, fibLocal)
		} else {
			break
		}
	}

	return minor4m
}

func evenp(n int) bool {
	if n % 2 == 0 {
		return true
	} else {
		return false
	}
}

func evenSlicep(s []int) []int {
	even := []int{}

	for _, i := range s {
		if evenp(i) {
			even = append(even, i)
		} else {
			continue
		}
	}

	return even
}
