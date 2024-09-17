package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	str := ""
	base := "0123456789abcdef"

	for i, v := range arr {
		z01.PrintRune(rune(base[v/16]))
		z01.PrintRune(rune(base[v%16]))

		if (i+1)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}

		if v >= 32 && v <= 126 {
			str += string(v)
		} else {
			str += string('.')
		}
	}
	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}