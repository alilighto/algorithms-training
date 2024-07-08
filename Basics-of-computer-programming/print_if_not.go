package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print(PrintIfNot("abcdefz"))
// 	fmt.Print(PrintIfNot("abc"))
// 	fmt.Print(PrintIfNot(""))
// 	fmt.Print(PrintIfNot("14"))
// }

func PrintIfNot(s string) string {
	if len(s) < 3 || s == ""{
		return "G\n"
}
return "invalid input\n"
}