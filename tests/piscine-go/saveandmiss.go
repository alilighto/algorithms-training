package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 || num > len(arg) {
		return string(arg)
	}
	str := ""
	for i := 0; i < len(arg); i++ {
		if i != 0 && i%num == 0 {
			i += num
			if i > len(arg)-1 {
				break
			}
		}
		if i != len(arg) {
			str += string(rune(arg[i]))
		}
	}
	return str
}