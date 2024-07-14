package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
// 	fmt.Print(LastWord(" lorem,ipsum "))
// 	fmt.Print(LastWord(" "))
// }

func LastWord(s string) string {
	x := len(s) - 1
	str := ""
	str1:= ""
	for i := x; i > 0; i-- {
		if s[i] == ' ' && i == x {
			x--
		}else {
			if s[i] == ' ' {
				break
			}
			str += string(s[i])
		}
	}

	for i := len(str)-1; i >= 0 ; i-- {
		str1 += string(str[i])
	}
	return str1 + "\n"
}