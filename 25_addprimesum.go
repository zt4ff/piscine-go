package piscine

import "github.com/01-edu/z01"

func atoi2(n string) int {
	result := 0
	for _, x := range n {
		result = result*10 + int(x-'0')
	}

	return result
}

func isPrimeAgain(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func printSing(s string) {
	for _, x := range s {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}

func AddPrimeSum(n string) {
	nInt := atoi2(n)
	count := 0

	for i := 2; i <= nInt; i++ {
		if isPrimeAgain(i) {
			count += i
		}
	}

	// convert count to string
	nStr := itoa(count)

	printSing(nStr)
}
