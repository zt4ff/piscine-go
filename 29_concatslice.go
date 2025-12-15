package piscine

func ConcatSlice(slice1, slice2 []int) []int {
	nSlice1 := len(slice1)
	nSlice2 := len(slice2)

	if nSlice1 == 0 {
		return slice2
	}

	if nSlice2 == 0 {
		return slice1
	}

	ans := []int{}

	for i := 0; i < nSlice1; i++ {
		ans = append(ans, slice1[i])
	}

	for i := 0; i < nSlice2; i++ {
		ans = append(ans, slice2[i])
	}

	return ans
}
