package piscine

func containOnlyAlphabet(str string) bool {
	for _, c := range str {
		if (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') {
			return false
		}
	}
	return true
}

func isUpper(s rune) bool {
	return s >= 'A' && s <= 'Z'
}

func CamelToSnakeCase(s string) string {
	res := ""
	if len(s) == 0 || !containOnlyAlphabet(s) {
		return s
	}
	for i := 0; i < len(s); i++ {
		if i != 0 && isUpper(rune(s[i])) && i+1 < len(s) && !isUpper(rune(s[i+1])) {
			res += "_"
			res += string(s[i])
		} else if !isUpper(rune(s[i])) || (i == 0 && isUpper(rune(s[i]))) {
			res += string(s[i])
		} else {
			return s
		}
	}
	return res
}