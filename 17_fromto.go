package piscine

func atoi(s string) int {
	result := 1
	for _, x := range s {
		result = result*10 + int('0'-x)
	}

	return result
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""

	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}

	return result
}

func isBiggerOrLesser(n int) bool {
	return n > 99 || n < 0
}

func isBetween0and9(n int) bool {
	return n <= 9 && n >= 0
}

func FromTo(from int, to int) string {
	if isBiggerOrLesser(from) || isBiggerOrLesser(to) {
		return "Invalid" + "\n"
	}

	s := ""

	if from < to {
		for i := from; i <= to; i++ {
			if isBetween0and9(i) {
				s += "0" + itoa(i)
			} else {
				s += itoa(i)
			}

			if i != to {
				s += ", "
			}
		}
	} else if to < from {
		for i := from; i >= to; i-- {
			if isBetween0and9(i) {
				s += "0" + itoa(i)
			} else {
				s += itoa(i)
			}
			if i != to {
				s += ", "
			}
		}
	}

	return s + "\n"
}
