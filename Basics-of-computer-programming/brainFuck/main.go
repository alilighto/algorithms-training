package main

import(
	"os"
)
func main() {
	args := os.Args[1:]
	if len(args) != 1 || len(args) > 4096 {
		return
	}
	arr := make([]byte, 2048)
	ptr := &arr[0]
	for v := range args{
		if v > 0{
			return ptr
		}
	}

}