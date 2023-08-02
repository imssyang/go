package dep

func LastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

func LastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}
