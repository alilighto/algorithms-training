package piscine

func RepeatAlpha(s string) string {
	str := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' ||
			v >= 'A' && v <= 'Z' {
			c := int(v - 'A')
			if int(v) >= 'a' {
				c = int(v - 'a')
			}
			for i := 0; i <= c; i++ {
				str += string(v)
			}
		} else {
			str += string(v)
		}
	}
	return str
}
