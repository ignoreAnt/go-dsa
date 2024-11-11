package gcd

// gcdIterative returns the greatest common divisor of a and b using the Euclidean
// algorithm in an iterative manner.
func gcdIterative(a, b int) int {

	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	// Euclidean algorithm
	//
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// gcdRecursive returns the greatest common divisor of a and b using the Euclidean
// algorithm in a recursive manner.
func gcdRecursive(a, b int) int {
	if b == 0 {
		return a
	}

	return gcdRecursive(b, a%b)
}
