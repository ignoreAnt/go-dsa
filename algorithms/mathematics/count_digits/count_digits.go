package count_digits

// CountDigits returns the number of digits in a given integer n.
//
// The function takes an int as argument and returns the number of digits in that int.
// If n is 0, the function returns 1.
func CountDigits(n int) int {

	if n == 0 {
		return 1
	}

	count := 0
	for n != 0 {
		n = n / 10 // n divided by 10 removes the last digit
		count++
	}

	return count
}
