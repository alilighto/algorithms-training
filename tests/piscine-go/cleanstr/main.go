package main

import "os"

func main(){
	c := false 
	str := ""
	args := os.Args[1:]
	if args[0] == "" || len(args) != 1 {
		os.Stdout.WriteString("\n")
		return
	} 
	for i, v := range args[0] {
		if v != ' ' {
			c = true
			str += string(v)
		}
		if v == ' ' && c && i < len(args[0])-1 && args[0][i+1] != ' '{
			str += " "
			c = false
		}
	}
	os.Stdout.WriteString(str+"\n")
}