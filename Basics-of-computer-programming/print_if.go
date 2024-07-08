package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print(PrintIf("abcdefz"))
// 	fmt.Print(PrintIf("abc"))
// 	fmt.Print(PrintIf(""))
// 	fmt.Print(PrintIf("14"))
// }

func PrintIf(s string) string {
	if len(s) > 3 || s == ""{
		return "G\n"
}
return "invalid input\n"
}