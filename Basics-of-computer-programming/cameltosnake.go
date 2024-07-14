package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(CamelToSnakeCase("HelloWorld"))
// 	fmt.Println(CamelToSnakeCase("helloworDD"))
// 	fmt.Println(CamelToSnakeCase("helloWorld"))
// 	fmt.Println(CamelToSnakeCase("camelCase"))
// 	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
// 	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
// 	fmt.Println(CamelToSnakeCase("hey2"))
// }

func CamelToSnakeCase(s string) string{
	if s == "" {
		return ""
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' && i != 0 {
			if i+1 <= (len(s)-1) && (s[i+1] >= 'A' && s[i+1] <= 'Z') || 
			s[i] < 'A' && s[i] > 'Z' {
				return s
			}
			s = s[:i] + "_" + s[i:]
			i++
		}
	}
	return s
}