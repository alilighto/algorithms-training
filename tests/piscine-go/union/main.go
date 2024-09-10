package main

import (
	"os"
)

func main() {
	if len(os.Args) == 3 {
		s1 := os.Args[1]
		s2 := os.Args[2]
		res := ""

		// Iterate through the first string
		for _, v := range s1 {
			if !containsRune(res, v) {
				res += string(v)
			}
		}

		// Iterate through the second string
		for _, v := range s2 {
			if !containsRune(res, v) {
				res += string(v)
			}
		}

		// Write the result to stdout
		os.Stdout.WriteString(res)
	}
	os.Stdout.WriteString("\n")
}

// Helper function to check if a rune is already in the result string
func containsRune(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
