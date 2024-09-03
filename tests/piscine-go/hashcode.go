package piscine

func HashCode(s string) string {
	str := ""
	for _, v := range s {
		s := ((int(v) + len(s)) % 127)
		if s < 32 || s > 126 {
			s += 33
		}
		str += string(rune(s))
	}
	return str
}
