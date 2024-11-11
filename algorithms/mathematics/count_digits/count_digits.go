package count_digits

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
