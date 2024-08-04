package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
// 	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
// }

// $ go run .
// [1 4 2 5 3 6]
// [1 2 3 4 5 6 7 8 9 10 11]
// [4 1 5 2 6 3 7 8 9]
// [1 2 3]

func ConcatAlternate(slice1, slice2 []int) []int {

	res := []int{}
	
	if len(slice1) < len(slice2) {
		slice1, slice2 = slice2, slice1
	}
	for i, v := range slice1 {
		res = append(res, v)
		if i < len(slice2) {
			res = append(res, slice2[i])
		}
	}
	return res
}
