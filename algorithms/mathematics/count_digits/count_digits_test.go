package count_digits

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// TestCountDigits tests the CountDigits function for various cases.
//
// The test cases cover both positive and negative numbers, single and multiple
// digits, edge cases with large numbers, and the smallest and largest two-digit
// numbers.
//
// The test also covers the maximum and minimum values for 32-bit and 64-bit
// signed integers, as well as single-digit negative numbers and multiple digit
// negative numbers.
func TestCountDigits(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"Single Digit(positive)", 5, 1},
		{"Single Digit(negative)", -5, 1},
		{"Single Digit(zero)", 0, 1},
		{"Single largest digit", 9, 1},
		{"Multiple Digits(positive)", 123, 3},
		{"Multiple Digits(negative)", -123, 3},
		{"Edge case of large numbers", 1000000000, 10}, // Powers of ten
		{"Smallest two-digit number", 10, 2},
		{"Largest two-digit number", 99, 2},

		{"Largest int for 32-bit", 2147483647, 10},   // Maximum int for 32-bit signed
		{"Smallest int for 32-bit", -2147483648, 10}, // Minimum int for 32-bit signed
		{"Single-digit negative", -5, 1},             // Treat negative numbers by absolute value
		{"Multiple digits negative", -12345, 5},      // Negative multi-digit input

		{"Max int for 64-bit", 9223372036854775807, 19},  // Maximum int for 64-bit signed
		{"Min int for 64-bit", -9223372036854775808, 19}, // Minimum int for 64-bit signed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := CountDigits(tc.input)

			// Use require to immediately stop the test if a critical check fails.
			require.NotNil(t, actual, "Result should not be nil")

			// Use assert to check if actual and expected values are equal
			assert.Equal(t, tc.expected, actual, "Expected: %d, Got: %d", tc.expected, actual)
		})
	}
}
