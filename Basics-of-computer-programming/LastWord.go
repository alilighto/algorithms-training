package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
// 	fmt.Print(LastWord(" lorem,ipsum "))
// 	fmt.Print(LastWord(" "))
// }

// func LastWord(s string) string {
// 	x := len(s) - 1
// 	str := ""

// 	// Trim trailing spaces
// 	for x >= 0 && s[x] == ' ' {
// 		x--
// 	}

// 	// Find the last word
// 	for i := x; i >= 0; i-- {
// 		if s[i] == ' ' {
// 			break
// 		}
// 		str = string(s[i]) + str
// 	}

// 	return str
// }


func LastWord(s string) string {
	x := len(s) - 1
	str := ""

	// Trim trailing spaces
	// for x >= 0 && s[x] == ' ' {
	// 	x--
	// }

	// Find the last word
	for i := x; i >= 0; i-- {
		if s[i] == ' ' && i == x {
			x--
		} else {
			if s[i] == ' ' {
				break
			}
			str = string(s[i]) + str
		}
	}

	return str + "\n"
}
