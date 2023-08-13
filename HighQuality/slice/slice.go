package slice

func GetLastBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

func GetLastByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])
	return result
}
