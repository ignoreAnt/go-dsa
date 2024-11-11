package prime_factors

// primeFactorsBrute returns a slice of all prime factors of n,
// using a brute-force method with a time complexity of O(n).
func primeFactorsBrute(n int) []int {
	var factors []int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}

	return factors
}

// primeFactors returns a slice of all prime factors of n.
//
// This function uses a hybrid method. It first divides out all the factors of 2,
// then checks the odd factors from 3 up to sqrt(n). If n is a prime number
// greater than 2, it is added to the slice of factors.
func primeFactors(n int) []int {
	var factors []int

	// Divide out the factor of 2
	for n%2 == 0 {
		factors = append(factors, 2)
		n = n / 2
	}

	// Check for the odd factors from 3 up to sqrt(n)
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}

	// If n is a prime number greater than 2 add it to factors
	if n > 2 {
		factors = append(factors, n)

	}

	return factors
}
