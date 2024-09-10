package piscine

func Itoa(n int) string {
	sign, res := "", ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		sign = "-"
		n = -n
	}
	for n > 0 {
		res = string(rune(n%10+'0')) + res
		n /= 10
	}
	return sign + res
}