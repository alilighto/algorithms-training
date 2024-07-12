package training

// import "fmt"

// func main() {
// 	fmt.Println(CheckBrackets("'(johndoe)'"))                         // OK
// 	fmt.Println(CheckBrackets("'([)]'"))                              // Error
// 	fmt.Println(CheckBrackets("'' '{[(0 + 0)(1 + 1)](3*(-1)){()}}'")) // OK
// }

func CheckBrackets(s string) string {
	stack := []rune{}

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return "Error"
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return "Error"
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return "Error"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "OK"
	}
	return "Error"
}
