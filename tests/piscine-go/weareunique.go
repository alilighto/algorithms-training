package piscine

import "strings"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	var array [127]int
	co := 0
	for i, v := range str1 + str2 {
		if array[v] == 0 && ((i < len(str1) && !strings.Contains(str2, string(v))) ||
		 (i >= len(str1) && !strings.Contains(str1, string(v)))) {
			array[v] = 1
			co++
		}
	}
	return co
}