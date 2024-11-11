package divisors_of_number

func divisorsOfNumber(n int) []int {
	var divisors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	return divisors
}

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
