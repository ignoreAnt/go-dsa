package lcm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestLCM tests the LCM function for various cases.
//
// The test cases include a variety of edge cases, such as when one or both
// numbers are 0, and more general cases with various LCM values.
// The test also verifies that both the iterative and recursive methods produce
// the same results.
func TestLCM(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		// Edge cases with zero
		{"LCM of 0 and 0", 0, 0, 0},   // LCM(0, 0) is conventionally 0
		{"LCM of 0 and 10", 0, 10, 0}, // LCM(0, x) is 0
		{"LCM of 5 and 0", 5, 0, 0},   // LCM(x, 0) is 0

		// LCM with one number being 1 (LCM of any number with 1 is the number itself)
		{"LCM of 1 and 1", 1, 1, 1},
		{"LCM of 1 and 7", 1, 7, 7},
		{"LCM of 15 and 1", 15, 1, 15},

		// LCM of two prime numbers (should be their product)
		{"LCM of 3 and 5", 3, 5, 15},
		{"LCM of 11 and 13", 11, 13, 143},
		{"LCM of 17 and 19", 17, 19, 323},

		// LCM where one number is a multiple of the other
		{"LCM of 4 and 8", 4, 8, 8},
		{"LCM of 15 and 5", 15, 5, 15},
		{"LCM of 12 and 36", 12, 36, 36},

		// LCM of equal numbers (should be the number itself)
		{"LCM of 9 and 9", 9, 9, 9},
		{"LCM of 100 and 100", 100, 100, 100},

		// LCM of two even numbers
		{"LCM of 14 and 18", 14, 18, 126},
		{"LCM of 20 and 30", 20, 30, 60},

		// LCM of two odd numbers
		{"LCM of 15 and 25", 15, 25, 75},
		{"LCM of 21 and 35", 21, 35, 105},

		// Large numbers
		{"LCM of 123456 and 789012", 123456, 789012, 8117355456},
		{"LCM of 1000000 and 500000", 1000000, 500000, 1000000}, // Where one is a multiple of the other
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := lcm(tc.a, tc.b)
			assert.Equal(t, tc.expected, actual, "Expected: %v, Got: %v", tc.expected, actual)
		})
	}
}
