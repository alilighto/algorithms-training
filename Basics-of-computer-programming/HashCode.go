package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(HashCode("A"))
// 	fmt.Println(HashCode("AB"))
// 	fmt.Println(HashCode("BAC"))
// 	fmt.Println(HashCode("Hello World"))
// }

func HashCode(s string) string {
	str := ""
	for _, v := range s {
		if v < 32 || v > 126 {
			str += string(v + 33)
		}
		str += string((v + rune(len(s)))%127)
	}
	return str
}

// package solutions

// func HashCode(dec string) string {
// 	size := len(dec)
// 	hashed := ""
// 	for _, char := range dec {
// 		hash := (int(char) + size) % 127
// 		if hash < 32 || hash > 126 {
// 			hash += 33
// 		}
// 		hashed += string(hash)
// 	}
// 	return hashed
// }