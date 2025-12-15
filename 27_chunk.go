package piscine

import "fmt"

func Chunk(slice []int, size int) {
	ans := [][]int{}

	if size <= 0 {
		fmt.Println(ans)
		return
	}

	for i := 0; i < len(slice); i += size {
		start := i
		end := i + size

		if end > len(slice) {
			end = len(slice)
		}

		ans = append(ans, slice[start:end])
	}

	fmt.Println(ans)
}
