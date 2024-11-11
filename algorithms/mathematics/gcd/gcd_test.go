package gcd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGCD tests the GCD function for various cases.
//
// The test cases include a variety of edge cases, such as when one or both
// numbers are 0, and more general cases with various GCD values.
// The test also verifies that both the iterative and recursive methods produce
// the same results.
func TestGCD(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		// Edge cases with zero
		{"GCD of 0 and 0", 0, 0, 0}, // Conventionally GCD(0, 0) can be considered 0 or undefined
		{"GCD of 0 and 5", 0, 5, 5}, // GCD(x, 0) should be x
		{"GCD of 7 and 0", 7, 0, 7}, // GCD(x, 0) should be x

		// GCD with one number being 1 (any number GCD with 1 is 1)
		{"GCD of 1 and 1", 1, 1, 1},
		{"GCD of 1 and 10", 1, 10, 1},
		{"GCD of 25 and 1", 25, 1, 1},

		// GCD of two prime numbers (should be 1)
		{"GCD of 13 and 17", 13, 17, 1},     // Both are primes
		{"GCD of 101 and 103", 101, 103, 1}, // Both are primes

		// GCD where one number is a multiple of the other
		{"GCD of 8 and 4", 8, 4, 4},
		{"GCD of 15 and 5", 15, 5, 5},
		{"GCD of 100 and 25", 100, 25, 25},

		// GCD of equal numbers (should be the number itself)
		{"GCD of 12 and 12", 12, 12, 12},
		{"GCD of 45 and 45", 45, 45, 45},

		// GCD of two even numbers
		{"GCD of 56 and 98", 56, 98, 14},
		{"GCD of 48 and 18", 48, 18, 6},

		// GCD of two odd numbers
		{"GCD of 15 and 35", 15, 35, 5},
		{"GCD of 63 and 21", 63, 21, 21},

		// Large numbers
		{"GCD of 270 and 192", 270, 192, 6},                                   // Common factors
		{"GCD of 123456 and 789012", 123456, 789012, 12},                      // Large values with common factors
		{"GCD of 1000000000 and 500000000", 1000000000, 500000000, 500000000}, // Large numbers where one is multiple of the other
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := gcdIterative(tc.a, tc.b) // Test the iterative method
			assert.Equal(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)

			actual = gcdRecursive(tc.a, tc.b) // Test the recursive method
			assert.Equal(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})
	}
}
