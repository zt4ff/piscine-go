package piscine

func Itoa(n int) string {
	neg := false

	if n < 0 {
		neg = true
		n *= -1
	}

	result := ""

	for n > 0 {
		d := n % 10
		result = string(d+'0') + result
		n /= 10
	}

	if neg {
		return "-" + result
	}

	return result
}
