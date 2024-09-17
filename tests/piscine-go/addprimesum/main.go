package main

import (
	"os"
)

func Atoi(s string) int {
	result, sign := 0, 1

	for i, char := range s {
		if i == 0 && (char == '-' || char == '+') {
			if char == '-' {
				sign = -1
			}
			continue
		}

		if char < '0' || char > '9' {
			return 0
		}

		digit := int(char - '0')
		result = result*10 + digit
	}
	return result * sign
}

func Itoaa(n int) string {
	sign, res := "", ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		sign = "-"
		n = -n
	}
	for n > 0 {
		res = string(rune(n%10+'0')) + res
		n /= 10
	}
	return sign + res
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumOfPrimes(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("0" + "\n")
		return
	}

	num := Atoi(os.Args[1])
	if num <= 0 {
		os.Stdout.WriteString("0" + "\n")
		return
	}
	os.Stdout.WriteString(Itoaa(sumOfPrimes(num)) + "\n")
}
