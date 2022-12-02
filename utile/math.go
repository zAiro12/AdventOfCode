package utile

// dati x numeri restituise il maggiore
func Max(a ...int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

// dati x numeri restituise il minore
func Min(a ...int) int {
	min := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > min {
			min = a[i]
		}
	}
	return min
}
