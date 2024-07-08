package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(CheckNumber("Hello"))
// 	fmt.Println(CheckNumber("Hello1"))
// }

func CheckNumber(s string) bool {
	for _, v := range s {
		if v >= '0' && v <= '9' {
			return true
		}
	}
	return false
}
