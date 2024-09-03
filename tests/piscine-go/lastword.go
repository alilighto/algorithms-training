package piscine

func LastWord(s string) string {
	c := false
	str := ""
	for i := len(s)-1; i >= 0; i-- {
		if s[i] != ' ' {
			c = true
			str = string(s[i]) + str
		} else if s[i] == ' ' && c {
			break
		}
	}
	return str + "\n"
}
