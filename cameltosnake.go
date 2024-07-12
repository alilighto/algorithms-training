package main

import (
	"fmt"
)

func main() {
	// fmt.Println(CamelToSnakeCase("HelloWorld"))
	// fmt.Println(CamelToSnakeCase("helloWorld"))
	// fmt.Println(CamelToSnakeCase("camelCase"))
	// fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	// fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string{
	st := ""
	for i, v := range s {
		if v >= 'A' && v <= 'Z' && i != 0 {
			i++
			st += "_" + string(s[i])
		}
	}
	return st
}