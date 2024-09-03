package piscine

func RectPerimeter(w, h int) int {
	if h < 0 || w < 0 {
		return -1
	}
	return 2*(h+w)
}