package piscine

func FirstWord(s string) string {
	c := false
	if s == "" {
		return "\n"
	}
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' '{
			c = true
			res += string(s[i])
		} else if s[i] == ' ' && c {
			break
		}
	}
	return res + "\n"
}
