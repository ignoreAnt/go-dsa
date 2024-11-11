package prime_factors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPrimeFactors tests the primeFactors function for various cases.
//
// The test cases include a variety of edge cases, small composite numbers,
// prime numbers, larger composite numbers with multiple prime factors, and
// numbers with higher powers of prime factors. The test also verifies that
// both the iterative (primeFactorsBrute) and recursive (primeFactors) methods
// produce the same results.
func TestPrimeFactors(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected []int
	}{
		// Edge cases
		{"Prime factors of 1", 1, nil},         // 1 has no prime factors
		{"Prime factors of 2", 2, []int{2}},    // 2 is prime
		{"Prime factors of 3", 3, []int{3}},    // 3 is prime
		{"Prime factors of 4", 4, []int{2, 2}}, // 4 = 2 * 2

		// Small composite numbers
		{"Prime factors of 6", 6, []int{2, 3}},    // 6 = 2 * 3
		{"Prime factors of 8", 8, []int{2, 2, 2}}, // 8 = 2 * 2 * 2
		{"Prime factors of 9", 9, []int{3, 3}},    // 9 = 3 * 3
		{"Prime factors of 10", 10, []int{2, 5}},  // 10 = 2 * 5

		// Prime numbers (should return the number itself as the only factor)
		{"Prime factors of 7", 7, []int{7}},    // 7 is prime
		{"Prime factors of 11", 11, []int{11}}, // 11 is prime
		{"Prime factors of 13", 13, []int{13}}, // 13 is prime
		{"Prime factors of 17", 17, []int{17}}, // 17 is prime

		// Larger composite numbers with multiple prime factors
		{"Prime factors of 18", 18, []int{2, 3, 3}}, // 18 = 2 * 3 * 3
		{"Prime factors of 20", 20, []int{2, 2, 5}}, // 20 = 2 * 2 * 5
		{"Prime factors of 28", 28, []int{2, 2, 7}}, // 28 = 2 * 2 * 7
		{"Prime factors of 30", 30, []int{2, 3, 5}}, // 30 = 2 * 3 * 5
		{"Prime factors of 50", 50, []int{2, 5, 5}}, // 50 = 2 * 5 * 5

		// Composite numbers with higher powers of prime factors
		{"Prime factors of 64", 64, []int{2, 2, 2, 2, 2, 2}}, // 64 = 2^6
		{"Prime factors of 81", 81, []int{3, 3, 3, 3}},       // 81 = 3^4
		{"Prime factors of 100", 100, []int{2, 2, 5, 5}},     // 100 = 2 * 2 * 5 * 5

		// Large prime number (should return the number itself)
		{"Prime factors of 97", 97, []int{97}}, // 97 is prime

		// Larger numbers with multiple prime factors
		{"Prime factors of 210", 210, []int{2, 3, 5, 7}},             // 210 = 2 * 3 * 5 * 7
		{"Prime factors of 231", 231, []int{3, 7, 11}},               // 231 = 3 * 7 * 11
		{"Prime factors of 256", 256, []int{2, 2, 2, 2, 2, 2, 2, 2}}, // 256 = 2^8
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			actual := primeFactorsBrute(tc.input)
			assert.Equal(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})

		t.Run(tc.name, func(t *testing.T) {
			actual := primeFactors(tc.input)
			assert.Equal(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})
	}
}
