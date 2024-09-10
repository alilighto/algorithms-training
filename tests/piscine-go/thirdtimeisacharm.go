package piscine

func ThirdTimeIsACharm(s string) string {
	str := ""
	co := 0
	for i := 0; i < len(s); i++ {
		if co == 2 {
			str += string(rune(s[i]))
			co = 0
		} else {
			co++
		}
	}
	return str + "\n"
}

// func ThirdTimeIsACharm(s string) string {
//     result := ""
//     for i := 2; i < len(s); i += 3 { // Start from the third character (index 2)
//         result += string(s[i]) // Append every third character to the result
//     }
//     return result
// }