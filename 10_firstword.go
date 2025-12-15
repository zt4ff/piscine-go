package piscine

func FirstWord(s string) string {
	// remove trailing spaces from s
	end := len(s) - 1
	for s[end] == ' ' {
		end--
	}

	if end == 0 {
		return "\n"
	}

	start := 0
	for s[start] == ' ' {
		start++
	}

	s = s[start : end+1]

	ans := ""

	for _, x := range s {
		if x == ' ' {
			break
		}
		ans += string(x)
	}

	ans += "\n"

	return ans
}
