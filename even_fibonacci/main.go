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
