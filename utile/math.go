package utile

import "math"

// dati x numeri restituise il maggiore
func Max[T Ordered](a ...T) T {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

// dati x numeri restituise il minore
func Min[T Ordered](a ...T) T {
	min := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}

// dato un numero restituisce il suo valore assoluto
func ABS[T Ordered](num T) T {
	return T(math.Abs(float64(num)))
}
