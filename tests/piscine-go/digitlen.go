package piscine

func DigitLen(a, b int) int {
	if b < 2 || b > 36 {
		return -1
	}
	if a < 0 {
		a = -a
	}
	if a == 0 {
		return 0
	}
	return 1 + DigitLen(a/b, b)
}