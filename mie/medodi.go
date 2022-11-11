package mie

func rimuoviInPos(input string, pos int) string {
	var out string
	out = input[:pos]
	out += input[pos+1:]
	return out
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
