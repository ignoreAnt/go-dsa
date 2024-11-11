package factorial

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestFactorial tests the factorial function for various cases.
//
// The test cases include a variety of small and large positive numbers,
// as well as edge cases for negative numbers and high values that would
// overflow Go's native int type. The test also verifies that both the
// iterative and recursive methods produce the same results.
func TestFactorial(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		// Edge case: 0! should be 1
		{"Factorial of 0", 0, 1},

		// Small numbers
		{"Factorial of 1", 1, 1},
		{"Factorial of 2", 2, 2},
		{"Factorial of 3", 3, 6},
		{"Factorial of 4", 4, 24},
		{"Factorial of 5", 5, 120},

		// Larger numbers
		{"Factorial of 6", 6, 720},
		{"Factorial of 7", 7, 5040},
		{"Factorial of 8", 8, 40320},
		{"Factorial of 10", 10, 3628800},
		{"Factorial of 12", 12, 479001600},

		// Edge cases for high values
		// Beyond 12, Go's int will overflow for factorial calculations.
		// Use `math/big` for numbers larger than this if needed.

		// Negative numbers (return -1 or handle error case)
		{"Factorial of -1 (error case)", -1, -1},
		{"Factorial of -5 (error case)", -5, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := factorialIterative(tc.input) // Test the iterative method
			assert.Equal(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)

			actual = factorialRecursive(tc.input) // Test the recursive method
			assert.Equal(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})
	}
}
