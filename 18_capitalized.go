package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	i := 0
	for i < len(s) {
		// Skip spaces
		for i < len(s) && s[i] == ' ' {
			i++
		}

		// If we reached the end after skipping spaces
		if i >= len(s) {
			break
		}

		// First character of the word
		c := s[i]

		// If it's a lowercase letter → fail
		if c >= 'a' && c <= 'z' {
			return false
		}

		// Move to the end of the word
		for i < len(s) && s[i] != ' ' {
			i++
		}
	}

	return true
}
