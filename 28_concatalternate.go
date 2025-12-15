package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	nSlice1 := len(slice1)
	nSlice2 := len(slice2)

	if nSlice1 == 0 {
		return slice2
	}

	if nSlice2 == 0 {
		return slice1
	}

	ans := []int{}

	if nSlice1 > nSlice2 {
		for i := 0; i < nSlice2; i++ {
			ans = append(ans, slice1[i])
			ans = append(ans, slice2[i])
		}
		ans = append(ans, slice1[nSlice2:nSlice1]...)
	} else if nSlice2 > nSlice1 {
		for i := 0; i < nSlice1; i++ {
			ans = append(ans, slice2[i])
			ans = append(ans, slice1[i])
		}
		ans = append(ans, slice2[nSlice1:nSlice2]...)
	} else {
		for i := 0; i < nSlice1; i++ {
			ans = append(ans, slice1[i])
			ans = append(ans, slice2[i])
		}
	}

	return ans
}
