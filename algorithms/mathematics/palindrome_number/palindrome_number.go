package palindrome_number

import "strconv"

// isPalindromeNumber checks if a given integer is a palindrome.
// A palindrome is a number that reads the same forward and backward.
// The function returns true if the integer is a palindrome, otherwise it returns false.
// Note: Negative numbers cannot be palindromic by definition.
func isPalindromeNumber(n int) bool {
	// Negative numbers can't be palindrome
	if n < 0 {
		return false
	}

	original, reversed := n, 0
	for n > 0 {
		lastDigit := n % 10 //
		reversed = reversed*10 + lastDigit
		n = n / 10
	}

	return original == reversed

}

// isPalindromeStringMethod checks if a given integer is a palindrome by
// converting the number into a string and comparing characters from both ends.
// A palindrome is a number that reads the same forward and backward.
// The function returns true if the integer is a palindrome, otherwise it returns false.
// Note: Negative numbers cannot be palindromic by definition.
func isPalindromeStringMethod(n int) bool {
	// Negative numbers can't be palindrome
	if n < 0 {
		return false
	}

	s := strconv.Itoa(n)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
