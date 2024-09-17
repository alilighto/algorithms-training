package piscine

func itoaa(n int) string {
	str, sign := "", ""
	if n < 0 {
		sign = "-"
		n = -n
	}
	if n == 0 {
		return "0"
	}
	for n > 0 {
		str = string(rune(n%10+48)) + str
		n /= 10
	}
	return sign + str
}

func ZipString(s string) string {
	str := ""
	co := 1
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == s[i+1] {
			co++
		} else {
			str += itoaa(co) + string(s[i])
			co = 1
		}
	}
	return str
}
