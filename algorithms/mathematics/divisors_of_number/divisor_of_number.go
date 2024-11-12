package divisors_of_number

// divisorsOfNumber returns a slice of all divisors of the input number.
//
// Example: divisorsOfNumber(12) returns [1, 2, 3, 4, 6, 12].
func divisorsOfNumber(n int) []int {
	var divisors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	return divisors
}

// allDivisorsOfNumber returns a slice of all divisors of the input number.
//
// This function runs in O(sqrt(n)) time, since it checks divisors up to sqrt(n).
// It is more efficient than divisorsOfNumber when n is large.
func allDivisorsOfNumber(n int) []int {
	var divisors []int

	// Check divisors up to sqrt(n)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}

	return divisors
}
