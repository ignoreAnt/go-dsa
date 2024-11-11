package factorial

func factorialIterative(n int) int {
	// Factorial is not defined for negative numbers
	if n < 0 {
		return -1
	}

	fact := 1
	for i := 2; i <= n; i++ {
		fact = fact * i
	}

	return fact
}

func factorialRecursive(n int) int {
	// Factorial is not defined for negative numbers
	if n < 0 {
		return -1
	}

	if n == 0 || n == 1 {
		return 1
	}

	return n * factorialRecursive(n-1)
}
