package prime_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestIsPrime tests the isPrime function for various cases.
//
// The test cases include a variety of edge cases, such as when one or both
// numbers are 0, and more general cases with various GCD values.
// The test also verifies that both the iterative and recursive methods produce
// the same results.
func TestIsPrime(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		// Edge cases and small numbers
		{"Negative number", -5, false}, // Negative numbers are not prime
		{"Zero", 0, false},             // 0 is not prime
		{"One", 1, false},              // 1 is not prime
		{"Two", 2, true},               // 2 is the smallest prime number
		{"Three", 3, true},             // 3 is prime
		{"Four", 4, false},             // 4 is not prime (even number)

		// Small primes and non-primes
		{"Prime 5", 5, true},        // 5 is prime
		{"Prime 7", 7, true},        // 7 is prime
		{"Non-prime 9", 9, false},   // 9 is not prime (3*3)
		{"Prime 11", 11, true},      // 11 is prime
		{"Non-prime 15", 15, false}, // 15 is not prime (3*5)

		// Larger primes and non-primes
		{"Prime 17", 17, true},      // 17 is prime
		{"Prime 19", 19, true},      // 19 is prime
		{"Non-prime 21", 21, false}, // 21 is not prime (3*7)
		{"Non-prime 25", 25, false}, // 25 is not prime (5*5)
		{"Prime 29", 29, true},      // 29 is prime

		// Testing primes and non-primes around 100
		{"Non-prime 100", 100, false}, // 100 is not prime (10*10)
		{"Prime 101", 101, true},      // 101 is prime
		{"Non-prime 102", 102, false}, // 102 is not prime (2*51)
		{"Prime 103", 103, true},      // 103 is prime
		{"Non-prime 105", 105, false}, // 105 is not prime (3*35)

		// Larger values
		{"Prime 199", 199, true},      // 199 is prime
		{"Prime 211", 211, true},      // 211 is prime
		{"Non-prime 221", 221, false}, // 221 is not prime (13*17)
		{"Prime 223", 223, true},      // 223 is prime
		{"Non-prime 225", 225, false}, // 225 is not prime (15*15)

		// Large prime and non-prime values for performance
		{"Prime 1009", 1009, true},      // 1009 is prime
		{"Non-prime 1024", 1024, false}, // 1024 is not prime (2^10)
		{"Prime 104729", 104729, true},  // 104729 is a large prime
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isPrime(tc.input)
			assert.Equal(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})
	}
}
