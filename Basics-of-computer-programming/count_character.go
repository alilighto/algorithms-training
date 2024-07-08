package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(CountChar("Hello World", 'l'))
// 	fmt.Println(CountChar("5  balloons", 5))
// 	fmt.Println(CountChar("   ", ' '))
// 	fmt.Println(CountChar("The 7 deadly sins", '7'))
// }

func CountChar(s string, c rune) int {
	count := 0
	for _, v := range s {
		if v == c {
			count++
		}
	}
	return count
}
