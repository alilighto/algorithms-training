package main

import (
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1, s2, res := os.Args[1], os.Args[2], ""
	seen := make(map[rune]bool)

	for _, char := range s1 {
		if !seen[char] && contains(s2, char) {
			res += string(char)
			seen[char] = true
		}
	}
	os.Stdout.WriteString(res + "\n")
}

func contains(s string, char rune) bool {
	for _, c := range s {
		if c == char {
			return true
		}
	}
	return false
}
