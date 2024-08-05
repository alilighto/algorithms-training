package main

import "fmt"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

// $ go run . | cat -e
// 68 65 6c 6c$
// 6f 10 15 2a$
// 00 00$
// hello..*..$
// $

func PrintMemory(b [10]byte) {
	for _, v := range b {
		fmt.Printf("%02x ", v)
	}
	fmt.Println()
}


// func PrintMemory(arr [10]byte) {
// 	for i := 0; i < len(arr); i += 4 {
// 		// Print hexadecimal representation
// 		for j := 0; j < 4 && i+j < len(arr); j++ {
// 			fmt.Printf("%02x ", arr[i+j])
// 		}
// 		fmt.Println()
// 	}

// 	// Print ASCII representation
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] >= 32 && arr[i] <= 126 {
// 			fmt.Printf("%c", arr[i])
// 		} else {
// 			fmt.Printf(".")
// 		}
// 	}
// 	fmt.Println()
// }