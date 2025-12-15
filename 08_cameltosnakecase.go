package piscine


func isCapital(r rune) bool {
	return r >= 'A' && r <= 'Z'
}


func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	seenCapital := false

	for _, x := range s {
		if isCapital(x) {
			if seenCapital {
				return s
			}
			seenCapital = true
		} else {
			seenCapital = false
		}
	}

	out := ""

	for i, x := range s {
		if isCapital(x) {
			if i != 0 {
				out += "_"
			}
		}
		out += string(x)
	}

	return out
}
