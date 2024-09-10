package main

import (
	"os"
)

func ok(s1 string, s2 string) bool {
	count := 0
	for i := 0; i < len(s1); i++ {
		found := false
		for j := count; j < len(s2); j++ {
			if s1[i] == s2[j] {
				count = j + 1
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) == 3 {
		if ok(os.Args[1], os.Args[2]) {
			os.Stdout.WriteString(os.Args[1] + "\n")
		}
	}
}

// package main

// import "os"

// func ok(s1, s2 string) bool {
// 	for i, j := 0, 0; i < len(s1); i++ {
// 		for j < len(s2) && s1[i] != s2[j] {
// 			j++
// 		}
// 		if j == len(s2) {
// 			return false
// 		}
// 		j++
// 	}
// 	return true
// }

// func main() {
// 	if len(os.Args) == 3 && ok(os.Args[1], os.Args[2]) {
// 		os.Stdout.WriteString(os.Args[1] + "\n")
// 	}
// }

