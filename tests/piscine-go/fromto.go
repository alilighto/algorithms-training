package piscine

func Itoascii(n int) string {
	str := ""
	i := ""
	if n < 0 {
		i += "-"
		n = -n
	}
	if n == 0 {
		return "0"
	}
	for n > 0 {
		str = string(rune('0' + n%10)) + str
		n /= 10
	}
	return str
}

func FromTo(start, end int) string {
	if start > 99 || end > 99 || start < 0 || end < 0 {
		return "Invalid\n"
	}

	if start == end {
		return Itoascii(start) + "\n"
	}

	str, step := "", 1
	if start > end {
		step = -1
	}

	for i := start; ; i += step {
		if i >= 0 && i <= 9 {
			str += "0"
		}
		str += Itoascii(i)
		if i == end {
			break
		}
		str += ", "
	}
	return str + "\n"
}