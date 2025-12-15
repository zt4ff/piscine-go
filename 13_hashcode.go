package piscine

func HashCode(dec string) string {
	res := ""
	size := len(dec)

	for _, x := range dec {
		hashed := (x + rune(size)) % 127
		if hashed < 32 {
			hashed = hashed + 33
		}

		// unprintable ascii is between 0 and 31
		res += string(hashed)
	}
	return res
}
