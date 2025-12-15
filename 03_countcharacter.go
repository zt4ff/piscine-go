package piscine

func CountChar(str string, c rune) int {
	count := 0
	for _, x := range str {
		if x == c {
			count++
		}
	}	

	return count
}