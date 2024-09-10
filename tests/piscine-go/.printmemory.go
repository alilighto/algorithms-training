// package piscine

// import "github.com/01-edu/z01"

// func PrintMemory(tab [10]byte) {
// 	str := ""
// 	for i, byt := range tab {
// 		// Convert byte to its hexadecimal representation
// 		printHex(byt)
// 		// Print space or newline as needed
// 		if (i+1)%4 == 0 || i == len(tab)-1 {
// 			z01.PrintRune('\n')
// 		} else {
// 			z01.PrintRune(' ')
// 		}
// 		// Build the string representation
// 		if byt >= 33 && byt <= 126 {
// 			str += string(byt)
// 		} else {
// 			str += "."
// 		}
// 	}
// 	printStr(str)
// 	z01.PrintRune('\n')
// }

// func printHex(b byte) {
// 	hex := "0123456789abcdef"
// 	z01.PrintRune(rune(hex[b/16]))
// 	z01.PrintRune(rune(hex[b%16]))
// }

// func printStr(s string) {
// 	for _, char := range s {
// 		z01.PrintRune(char)
// 	}
// }

package piscine

import "github.com/01-edu/z01"

func PrintMemory(tab [10]byte) {
	str := ""
	hex := "0123456789abcdef"

	for i, byt := range tab {
		// Print hexadecimal representation
		z01.PrintRune(rune(hex[byt/16]))
		z01.PrintRune(rune(hex[byt%16]))

		// Print space or newline
		if (i+1)%4 == 0 || i == len(tab)-1 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}

		// Build string for printable characters
		if byt >= 33 && byt <= 126 {
			str += string(byt)
		} else {
			str += "."
		}
	}

	// Print the final string
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
