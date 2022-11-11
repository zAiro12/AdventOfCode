package mie

type interi int

func rimuoviInPos(input string, pos int) string {
	var out string
	out = input[:pos]
	out += input[pos+1:]
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (x interi) isMin(y int) bool {

	if int(x) < y {
		return true
	} else {
		return false
	}
}
