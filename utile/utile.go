package utile

type coordinate struct {
	x, y int
}

func rimuoviElemento(s []int, i int) []int {
	if len(s) <= i {
		return s
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
