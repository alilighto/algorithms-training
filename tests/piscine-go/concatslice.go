package piscine

func ConcatSlice(a, b []int) []int {
	if len(a) == 0 && len (b) == 0 {
		return []int(nil)
	}
	return append(a, b...)
}