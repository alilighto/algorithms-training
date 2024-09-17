package main

import (
	"os"
)

func ToUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A') // Convert to uppercase by adjusting ASCII values
	}
	return r
}

func ToLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A') // Convert to lowercase by adjusting ASCII values
	}
	return r
}
func main() {
	for _, arg := range os.Args[1:] {
		arg := []rune(arg)
		for j, r := range arg {
			if j+1 == len(arg) || arg[j+1] == ' ' {
				arg[j] = ToUpper(r)
			} else {
				arg[j] = ToLower(r)
			}
		}
		os.Stdout.WriteString(string(arg) + "\n")
	}
}
