package piscine

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""
	last := rune(s[0])
	count := 1

	for _, c := range s[1:] {
		if c == last {
			count++
		} else {
			result += string('0'+count) + string(last)
			last = c
			count = 1
		}
	}

	result += string('0'+count) + string(last)

	return result
}
