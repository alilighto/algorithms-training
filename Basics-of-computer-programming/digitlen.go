package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(DigitLen(100, 10))
// 	fmt.Println(DigitLen(100, 2))
// 	fmt.Println(DigitLen(-100, 16))
// 	fmt.Println(DigitLen(100, -1))
// }

func DigitLen(a, b int) int {
	if a == 0 {
		return 0
	}
	if b < 2 || b > 36 {
		return -1
	}
	return 1 + DigitLen(a/b, b)
}