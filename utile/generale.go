package utile

type point struct {
	x, y int
}

func RimuoviElemento(s []int, index int) []int {
	a := s[:index]
	a = append(a, s[index+1:]...)
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
