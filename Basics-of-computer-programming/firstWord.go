package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print(FirstWord("hello there"))
// 	fmt.Print(FirstWord(""))
// 	fmt.Print(FirstWord("hello   .........  bye"))
// }

func FirstWord(s string) string {
	str := ""
	if s == "" {
		str += "\n"
	}
	for _, v := range s {
		if v == ' ' {
			str += "\n"
			break
		}
		str += string(v)
	}
	return str
}

// package solutions

// import "strings"

// func FirstWord(s string) string {
// 	words := strings.Fields(s)
// 	res := "\n"
// 	if len(words) > 0 {
// 		res = words[0] + res
// 	}
// 	return res
// }