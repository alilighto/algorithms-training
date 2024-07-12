package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
// 	fmt.Println(RetainFirstHalf("A"))
// 	fmt.Println(RetainFirstHalf(""))
// 	fmt.Println(RetainFirstHalf("Hello World"))
// }

func RetainFirstHalf(s string) string {
	if s == ""{
		return ""
	}
	if len(s) == 1 {
		return s
	}
	return s[:len(s)/2]
}