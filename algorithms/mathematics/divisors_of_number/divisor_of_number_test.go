package divisors_of_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllDivisors(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected []int
	}{
		// Edge cases and small numbers
		{"Divisors of 1", 1, []int{1}},    // Only 1 is a divisor of itself
		{"Divisors of 2", 2, []int{1, 2}}, // Divisors of 2 are 1 and 2
		{"Divisors of 3", 3, []int{1, 3}}, // Divisors of 3 are 1 and 3 (prime number)

		// Composite numbers with small divisors
		{"Divisors of 4", 4, []int{1, 4, 2}},    // Divisors of 4 are 1, 2, and 4
		{"Divisors of 6", 6, []int{1, 6, 2, 3}}, // Divisors of 6 are 1, 2, 3, and 6
		{"Divisors of 8", 8, []int{1, 8, 2, 4}}, // Divisors of 8 are 1, 2, 4, and 8

		// Perfect squares
		{"Divisors of 9", 9, []int{1, 9, 3}},          // Divisors of 9 are 1, 3, and 9
		{"Divisors of 16", 16, []int{1, 16, 2, 8, 4}}, // Divisors of 16 are 1, 2, 4, 8, and 16
		{"Divisors of 25", 25, []int{1, 25, 5}},       // Divisors of 25 are 1, 5, and 25

		// Larger prime number
		{"Divisors of 29", 29, []int{1, 29}}, // 29 is a prime number, so divisors are 1 and 29

		// Larger composite numbers
		{"Divisors of 30", 30, []int{1, 30, 2, 15, 3, 10, 5, 6}},         // Divisors of 30
		{"Divisors of 36", 36, []int{1, 36, 2, 18, 3, 12, 4, 9, 6}},      // Divisors of 36
		{"Divisors of 50", 50, []int{1, 50, 2, 25, 5, 10}},               // Divisors of 50
		{"Divisors of 100", 100, []int{1, 100, 2, 50, 4, 25, 5, 20, 10}}, // Divisors of 100

		// Large number for performance testing
		{"Divisors of 101", 101, []int{1, 101}},     // 101 is a prime, so divisors are 1 and 101
		{"Divisors of 121", 121, []int{1, 121, 11}}, // 121 is 11^2, so divisors are 1, 11, and 121
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := divisorsOfNumber(tc.input)
			assert.ElementsMatch(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})

		t.Run(tc.name, func(t *testing.T) {
			actual := allDivisorsOfNumber(tc.input)
			assert.ElementsMatch(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})
	}
}
