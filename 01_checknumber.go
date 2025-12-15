package piscine

func CheckNumber(arg string) bool {
	for _, x := range arg {
		if x >= '0' && x <= '9' {
			return true
		}
	}
	return false
}