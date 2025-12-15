package piscine

func DigitLen(n, base int) int {
	if !(base >= 2 && base <= 36) {return -1}

	// convert negative to positive
	if n < 0 {
		n = n * -1
	}

	count := 0

	for n / base >= 0 {
		if n == 0 {
			break
		}
		count++
		n = n / base
	}

	return count
}