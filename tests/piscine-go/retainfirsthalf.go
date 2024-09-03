package piscine

func RetainFirstHalf(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	return s[:len(s)/2]
}