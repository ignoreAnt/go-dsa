package palindrome_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestIsPalindrome tests the isPalindrome function for various cases.
func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		// Single-digit numbers (all are palindromic)
		{"Single digit (0)", 0, true},
		{"Single digit (1)", 1, true},
		{"Single digit (9)", 9, true},

		// Small positive palindromic numbers
		{"Two digits, palindromic", 11, true},
		{"Three digits, palindromic", 121, true},
		{"Four digits, palindromic", 1221, true},
		{"Five digits, palindromic", 12321, true},

		// Small non-palindromic numbers
		{"Two digits, non-palindromic", 10, false},
		{"Three digits, non-palindromic", 123, false},
		{"Four digits, non-palindromic", 1234, false},
		{"Five digits, non-palindromic", 12345, false},

		// Large palindromic numbers
		{"Six digits, palindromic", 100001, true},
		{"Seven digits, palindromic", 1234321, true},
		{"Eight digits, palindromic", 12211221, true},
		{"Nine digits, palindromic", 123454321, true},
		{"Ten digits, palindromic", 1000000001, true},

		// Large non-palindromic numbers
		{"Six digits, non-palindromic", 123456, false},
		{"Seven digits, non-palindromic", 1234567, false},
		{"Eight digits, non-palindromic", 12345678, false},
		{"Nine digits, non-palindromic", 123456789, false},

		// Edge cases with maximum integer values for 32-bit and 64-bit
		{"Max int for 32-bit, non-palindromic", 2147483647, false},
		{"Max int for 64-bit, non-palindromic", 9223372036854775807, false},
		{"Nine nines, palindromic", 999999999, true}, // Large palindrome within 32-bit range

		// Edge cases with negative numbers (all should be false as negatives are non-palindromic by definition)
		{"Negative single digit", -1, false},
		{"Negative two digits", -11, false},
		{"Negative three digits", -121, false},
		{"Negative max int for 32-bit", -2147483648, false},
		{"Negative max int for 64-bit", -9223372036854775808, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isPalindromeNumber(tc.input)
			assert.Equal(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})

		t.Run(tc.name, func(t *testing.T) {
			actual := isPalindromeStringMethod(tc.input)
			assert.Equal(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})
	}
}
