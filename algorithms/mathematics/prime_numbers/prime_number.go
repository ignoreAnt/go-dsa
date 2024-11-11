package prime_numbers

// isPrime checks if a given integer is a prime number.
//
// A prime number is a natural number greater than 1 that has no positive divisors other than 1 and itself.
// The function returns true if the integer is a prime number, otherwise it returns false.
func isPrime(n int) bool {

	// Negative numbers and 1 are not prime
	if n <= 1 {
		return false
	}

	// 2, 3, and 5 are prime
	if n <= 3 {
		return true
	}

	// Check if n is divisible by 2 or 3
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Check if n is divisible by 6k + 1 or 6k + 5
	// for k = 1, 2, 3, ...
	// Check all odd numbers from 5 to sqrt(n)
	//
	// Efficient solution :
	//      Idea : Divisors always appear in pairs
	//      30 : (30,1) ; (2, 15); (3, 10) ; (5, 6)
	//      65 : (1, 65) ; (5, 13)
	//      25 : (1, 25) ; (5, 5)
	//
	//      if (x, y) is a Pair
	//      ==> x * y = n
	//      and if x <= y
	//      ==> x * x <= n
	//      ==> x <= sqrt(n)
	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
