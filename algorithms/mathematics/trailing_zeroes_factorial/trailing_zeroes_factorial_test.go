package trailing_zeroes_factorial

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestTrailingZeroes tests the TrailingZeroesInFactorial function for various cases.
//
// This test covers a range of values, including base cases, cases with 1 trailing zero,
// cases with multiple trailing zeros, and larger numbers to check the efficiency of the
// algorithm. It also tests the limits of the function by giving very large values.
func TestTrailingZeroes(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		// Base cases
		{"Factorial of 0", 0, 0}, // 0! = 1 (no trailing zeros)
		{"Factorial of 1", 1, 0}, // 1! = 1 (no trailing zeros)
		{"Factorial of 2", 2, 0}, // 2! = 2 (no trailing zeros)
		{"Factorial of 3", 3, 0}, // 3! = 6 (no trailing zeros)
		{"Factorial of 4", 4, 0}, // 4! = 24 (no trailing zeros)

		// Cases with 1 trailing zero
		{"Factorial of 5", 5, 1}, // 5! = 120 (1 trailing zero)
		{"Factorial of 6", 6, 1}, // 6! = 720 (1 trailing zero)
		{"Factorial of 9", 9, 1}, // 9! = 362880 (1 trailing zero)

		// Cases with multiple trailing zeros
		{"Factorial of 10", 10, 2}, // 10! = 3628800 (2 trailing zeros)
		{"Factorial of 15", 15, 3}, // 15! has 3 trailing zeros
		{"Factorial of 20", 20, 4}, // 20! has 4 trailing zeros
		{"Factorial of 25", 25, 6}, // 25! has 6 trailing zeros

		// Larger numbers to check efficiency
		{"Factorial of 50", 50, 12},      // 50! has 12 trailing zeros
		{"Factorial of 100", 100, 24},    // 100! has 24 trailing zeros
		{"Factorial of 125", 125, 31},    // 125! has 31 trailing zeros
		{"Factorial of 200", 200, 49},    // 200! has 49 trailing zeros
		{"Factorial of 1000", 1000, 249}, // 1000! has 249 trailing zeros

		// Very large values to test limits
		{"Factorial of 5000", 5000, 1249},
		{"Factorial of 10000", 10000, 2499},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := TrailingZeroesInFactorial(tc.input)
			assert.Equal(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})
	}
}
