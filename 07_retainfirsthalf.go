package piscine

func RetainFirstHalf(str string) string {
	n := len(str)

	if n == 0 || n == 1 {return str}

	return str[:n/2]
}