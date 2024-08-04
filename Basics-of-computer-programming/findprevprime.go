package training

// import "fmt"

// func main() {
// 	fmt.Println(FindPrevPrime(5000000000))
// 	fmt.Println(FindPrevPrime(4))
// }

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}


func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	if IsPrime(nb) {
		return nb
	}
	return FindPrevPrime(nb - 1)
}