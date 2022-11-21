package utile

func rimuoviElemento(s []int, index int) []int {
	a := s[:index]
	a = append(a, s[index+1:]...)
	return a
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
