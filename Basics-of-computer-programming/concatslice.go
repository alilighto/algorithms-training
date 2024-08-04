package training

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
// }

// $ go run .
// [1 2 3 4 5 6]
// [4 5 6 7 8 9]
// [1 2 3]

func ConcatSlice(s1, s2 []int) []int {
    return append(s1, s2...)
}