package piscine

func isPrime(n int) bool {
	// 0, 1 and 2 are not prime number
	if n < 2 {
		return false
	}

	// check divisions from 2 up until i * i <= n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func FindPrevPrime(nb int) int {
	for i := nb; i >= 3; i-- {
		if isPrime(i) {
			return i
		}
	}

	return 0
}
