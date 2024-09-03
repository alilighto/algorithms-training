package piscine

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 { 
		return 0
	}
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}
