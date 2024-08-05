package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print(FromTo(1, 10))
// 	fmt.Print(FromTo(10, 1))
// 	fmt.Print(FromTo(10, 10))
// 	fmt.Print(FromTo(100, 10))
// }

func itoa(a int) string {
	str := ""
	i := ""
	if a == 0 {
		return "0"
	}
	if a < 0 {
		i = "-"
	}
	for a > 0 {
		str = string('0' + a%10) + str
		a /= 10
	}
	return i + str
}

// soumi
// asoudnk

func FromTo(start, end int) string {
	str := ""
	if start > 99 || end > 99 {
        return "Invalid\n"
    }
	if start < end { 
		for i := start; i <= end; i++ {
			if i >= 0 && i <= 9{
				str += "0"
			}
			str += itoa(i)
			if i == end {
				break
			}
			str += ", "
		}
	}
	if start > end {
		for i := start; i >= end; i-- {
			if i >= 0 && i <= 9{
				str += "0"
			}
			str += itoa(i)
			if i == end {
				break
			}
			str += ", "
		}
	}
	if start == end {
		str += Itoa(start)
	}
	return str + "\n"
}