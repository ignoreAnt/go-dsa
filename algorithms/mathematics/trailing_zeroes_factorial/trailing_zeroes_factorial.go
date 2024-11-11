package trailing_zeroes_factorial

// TrailingZeroesInFactorial returns the number of trailing zeroes in n factorial.
//
// A trailing zero is produced by a pair of 2 and 5. So we just need to count
// the number of 5's, since there are always enough 2's.
func TrailingZeroesInFactorial(n int) int {
	count := 0
	for i := 5; i <= n; i = i * 5 {
		count = count + n/i
	}
	return count
}
