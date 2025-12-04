package piscine

func HashCode(dec string) string {
	// unpritable is between 0 and 32 ascii code

	result := ""
	l := len(dec)

	for _, i := range dec {
		c := (int(i) + l) % 127
		if c < 32 {
			c = c + 33
		}

		result += string(c)
	}

	return result
}
