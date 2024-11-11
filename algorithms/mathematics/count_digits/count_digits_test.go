package count_digits

import "testing"

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
			got := CountDigits(tc.input)
			if got != tc.expected {
				t.Errorf("got %d, want %d", got, tc.expected)
			}
		})
	}
}
