package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(IsCapitalized("Hello! How are you?"))
// 	fmt.Println(IsCapitalized("Hello How Are You"))
// 	fmt.Println(IsCapitalized("Whats 4this 100K?"))
// 	fmt.Println(IsCapitalized("Whatsthis4"))
// 	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
// 	fmt.Println(IsCapitalized(""))
// }

// $ go run .
// false
// true
// true
// true
// true
// false

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' && i != 0 && s[i-1] == ' ' {
			return false
		}
	}
	return !(s[0] >= 'a' && s[0] <= 'z')
}
