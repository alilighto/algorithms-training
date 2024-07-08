package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(CountAlpha("Hello world"))
// 	fmt.Println(CountAlpha("H e l l o"))
// 	fmt.Println(CountAlpha("H1e2l3l4o"))
// }

func CountAlpha(s string) int {
	co := 0
	for _, v := range s {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			co++
		}
	}
	return co
}
