package main

func mult(n, mult int) bool {
	if n % mult == 0 {
		return true
	} else {
		return false
	}
}

func sumM(n int) int {
	cont := 0

	for i := 0; i < n; i++ {
		if mult(i, 3) || mult(i, 5) {
			cont += i
		}
	}

	return cont
}
