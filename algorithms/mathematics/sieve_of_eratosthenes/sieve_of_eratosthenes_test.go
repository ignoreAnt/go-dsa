package sieve_of_eratosthenes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSieveOfEratosthenes tests the primesBrute and sieveOfEratosthenes functions
// for various edge cases and typical cases.
//
// The test cases include a variety of edge cases, such as when one or both
// numbers are 0, and more general cases with various primes.
// The test also verifies that both the iterative (primesBrute) and recursive
// (sieveOfEratosthenes) methods produce the same results.
func TestSieveOfEratosthenes(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected []int
	}{
		// Edge cases
		{"Primes up to 0", 0, []int{}},     // No primes less than or equal to 0
		{"Primes up to 1", 1, []int{}},     // No primes less than or equal to 1
		{"Primes up to 2", 2, []int{2}},    // Only 2 is prime
		{"Primes up to 3", 3, []int{2, 3}}, // Primes up to 3 are 2 and 3

		// Small values
		{"Primes up to 5", 5, []int{2, 3, 5}},              // Primes up to 5 are 2, 3, 5
		{"Primes up to 10", 10, []int{2, 3, 5, 7}},         // Primes up to 10 are 2, 3, 5, 7
		{"Primes up to 15", 15, []int{2, 3, 5, 7, 11, 13}}, // Primes up to 15

		// Typical cases
		{"Primes up to 20", 20, []int{2, 3, 5, 7, 11, 13, 17, 19}},         // Primes up to 20
		{"Primes up to 30", 30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}}, // Primes up to 30

		// Larger values
		{"Primes up to 50", 50, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}},                                                                    // Primes up to 50
		{"Primes up to 100", 100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},                          // Primes up to 100
		{"Primes up to 120", 120, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}}, // Primes up to 120
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			actual := primesBrute(tc.input)
			assert.ElementsMatch(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})

		t.Run(tc.name, func(t *testing.T) {
			actual := sieveOfEratosthenes(tc.input)
			assert.ElementsMatch(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})
	}
}
