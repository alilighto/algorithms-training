package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(RepeatAlpha("abc"))
// 	fmt.Println(RepeatAlpha("Choumi."))
// 	fmt.Println(RepeatAlpha(""))
// 	fmt.Println(RepeatAlpha("abacadaba 01!"))
// }

func RepeatAlpha(s string) string {
	var str string
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			count := int(r - 'A')
			if r >= 'a' {
				count = int(r - 'a')
			}
			for j := 0; j <= count; j++ {
				str += string(r)
			}
		} else {
			str += string(r)
		}
	}
	return str
}

// func RepeatAlpha(s string) string {
// 	str := ""
// 	x := 0
// 	for _, r := range s {

// 		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
// 			if r >= 'A' && r <= 'Z' {
// 				x = int(r - 'A')
// 			}
// 			if r >= 'a' && r <= 'z' {
// 				x = int(r - 'a')
// 			}
// 			for j := 0; j <= x; j++{
// 				str += string(r)
// 			}
// 		} else {
// 			str += string(r)
// 		}
// 	}
// 	return str
// }