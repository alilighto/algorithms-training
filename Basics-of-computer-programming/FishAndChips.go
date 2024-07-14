package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(FishAndChips(4))
// 	fmt.Println(FishAndChips(9))
// 	fmt.Println(FishAndChips(6))
// }

func FishAndChips(a int) string {
	if a < 0 {
		return "error: number is negative"
	}
	if a%2 == 0 && a%3 == 0 {
		return "Fish and Chips"
	} else if a%3 == 0 {
		return "Chips"
	} else if a%2 == 0 {
		return "Fish"
	} else {
		return "error: non divisible"
	}
}
