package palindrome_number

import "strconv"

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
