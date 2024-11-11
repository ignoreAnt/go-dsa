package lcm

// gcd returns the greatest common divisor of a and b.
//
// The formula for the greatest common divisor is gcd(a, b) = gcd(b, a % b).
func gcd(a, b int) int {
	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// lcm returns the lowest common multiple of two integers.
//
// The formula for the lowest common multiple is `lcm(a, b) = (a * b) / gcd(a, b)`.
func lcm(a, b int) int {

	// Check if either a or b is 0,
	if a == 0 || b == 0 {
		return 0
	}

	return (a * b) / gcd(a, b)
}
