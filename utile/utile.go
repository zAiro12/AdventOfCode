package utile

func rimuoviElemento(s []int, i int) []int {
	if len(s) <= i {
		return s
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}