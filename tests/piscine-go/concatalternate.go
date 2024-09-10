package piscine

func ConcatAlternate(a, b []int) []int {
	if len(b) > len(a) {
		a, b = b, a
	}
	var res []int
	for i := 0; i < len(a); i++ {
		res = append(res, a[i])
		if i < len(b) {
			res = append(res, b[i])
		}
	}
	return res
}
