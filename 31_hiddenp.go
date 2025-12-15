package piscine

import (
	"fmt"
	"os"
)

func contains1(s string, c rune) int {
	ans := -1
	for i, x := range s {
		if x == c {
			ans = i
		}
	}
	return ans
}

func HiddenP() {
	if len(os.Args) != 3 {
		return
	}

	s1 := (os.Args[1])
	s2 := (os.Args[2])

	lastIndex := 0

	for _, x := range s1 {
		n := contains1(s2, x)
		if n < 0 {
			fmt.Println("0")
			return
		}
		if lastIndex <= n {
			lastIndex = n
			continue
		} else {
			fmt.Println("0")
			return
		}
	}

	fmt.Println("1")
}
