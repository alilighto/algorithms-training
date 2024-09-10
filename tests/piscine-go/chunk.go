package piscine

import "fmt"

func Chunk(a []int, ch int) {
	if ch <= 0 {
		fmt.Println()
		return
	}
	var slice []int
	res := [][]int{}
	for len(a) >= ch {
		slice, a = a[:ch], a[ch:]
		res = append(res, slice)
	}
	if len(a) > 0 {
		res = append(res, a)
	}
	fmt.Println(res)
}