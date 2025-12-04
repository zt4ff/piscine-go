package piscine

func Gcd(a, b uint) uint {
	var result uint = 1
	if b < a {
		for i := b; i >= 1; i-- {
			if a%i == 0 && b%i == 0 {
				return i
			}
		}
	} else {
		for i := a; i >= 1; i-- {
			if a%i == 0 && b%i == 0 {
				return i
			}
		}
	}

	return result
}
