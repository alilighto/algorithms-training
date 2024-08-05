package training

// import (
// 	"fmt"
// )

// func main() {
//     fmt.Println(Itoa(12345))
//     fmt.Println(Itoa(0))
//     fmt.Println(Itoa(-1234))
//     fmt.Println(Itoa(987654321))
// }

func Itoa(n int) string {
	str := ""
	i := ""
	if n < 0 {
        i += "-"
        n = -n
    }
	if n == 0 {
        return "0"
    }
	for n > 0 {
        str = string('0' + n%10) + str
        n /= 10
    }
	return i + str
}