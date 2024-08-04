package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print(WordFlip("   First second last        "))
// 	fmt.Print(WordFlip(""))
// 	fmt.Print(WordFlip("     "))
// 	fmt.Print(WordFlip(" hello  all  of  you! "))
// }

// $ go run . | cat -e
// last second First$
// Invalid Output$
// $
// you! of all hello$

func Split(s string) []string {
	words := []string{}
	sl := ""
	for i, v := range s {
		if v != ' ' {
			sl += string(v)
		}
		if v == ' ' || i == len(s)-1 {
			words = append(words, sl)
			sl = ""
		}
	}
	return words
}

func TrimSpaces(s string) string {
	str := ""
	inWord := false
	for _, v := range s {
		if v != ' ' {
			str += string(v)
			inWord = true
		} else if inWord {
			str += " "
			inWord = false
		}
	}
	if str == "" {
		return ""
	}
	if str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}
	return str
}

// func WordFlip(str string) string {
// 	slice := []string{}
// 	res := ""
// 	if str == "" {
// 		return "Invalid Output\n"
// 	}

// 	str = TrimSpaces(str)
// 	if str == "" {
// 		return "\n"
// 	}
// 	words := Split(str)
// 	slice = append(slice, words[len(words)-1])
// 	slice = append(slice, words[1:len(words)-1]...)
// 	slice = append(slice, words[0])
// 	for _, v := range slice {
// 		res += v + " "
// 	}
// 	return res + "\n"
// }


func WordFlip(arg string) string {
	if arg == "" {
		return "Invalid Output\n"
	}
	str := Split(arg)
	res := ""
	for i := len(str)-1; i >= 0; i-- {
		if len(str[i]) != 0 {
			res += str[i]
		}
		if i > 0 && len(str[i-1]) != 0 {
			res += " "
		}
	}
	return (TrimSpaces(res) + "\n")
}