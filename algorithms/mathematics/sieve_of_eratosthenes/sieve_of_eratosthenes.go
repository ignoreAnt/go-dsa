package sieve_of_eratosthenes

// primesBrute returns a slice of all prime numbers from 2 up to n inclusive.
// This function uses a brute-force method with a time complexity of O(n).
func primesBrute(n int) []int {
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// isPrime checks if a given integer is a prime number.
//
// A prime number is a natural number greater than 1 that has no positive divisors other than 1 and itself.
// The function returns true if the integer is a prime number, otherwise it returns false.
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// sieveOfEratosthenes returns a slice of all prime numbers from 2 up to n inclusive.
// This function uses the Sieve of Eratosthenes algorithm with a time complexity of O(n).
// It iterates over the numbers from 2 to n and marks the multiples of each prime as non-prime.
// Finally, it returns the list of all remaining prime numbers.
func sieveOfEratosthenes(n int) []int {
	if n < 2 {
		return []int{}
	}

	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
