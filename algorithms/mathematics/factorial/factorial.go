package factorial

// factorialIterative calculates the factorial of n (n!) using an iterative approach.
//
// If n is negative, the function returns -1, as the factorial of a negative number
// is undefined.
//
// Otherwise, it initializes a variable, fact, to 1, and then iterates from 2 to n,
// multiplying fact by i in each iteration.
//
// Finally, it returns the calculated factorial.
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

// factorialRecursive calculates the factorial of a non-negative integer n using recursion.
//
// If n is negative, the function returns -1, as the factorial of a negative number
// is undefined.
//
// For non-negative integers, the function recursively multiplies n by the factorial
// of (n-1), with the base case being that both 0! and 1! are equal to 1.
//
// Parameters:
//
//	n int - the integer for which to calculate the factorial.
//
// Returns:
//
//	int - the factorial of n, or -1 if n is negative.
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
