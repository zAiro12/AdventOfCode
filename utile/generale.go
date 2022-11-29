package utile

import (
	"sort"
	"strings"
)

//struct con coordinate x, y
type Point struct {
	X, Y int
}

type Interfaccia interface{
	//sort.Interface
}

// data una slice di interi elimina l'elemento in posizione x
func RimuoviElementoInt(arr []int, index int) []int {
	a := arr[:index]
	a = append(a, arr[index+1:]...)
	return a
}

// data una slice di string elimina l'elemento in posizione x
func RimuoviElementoString(arr []string, index int) []string {
	a := arr[:index]
	a = append(a, arr[index+1:]...)
	return a
}

// data una slice di qualsiasi tipo elimina l'elemento in posizione x
func RimuoviElemento(arr []any, index int)[]any{
	a := arr[:index]
	a = append(a, arr[index+1:]...)
	return a
}

// dati x numeri restituise il maggiore
func Max(a ...int) int {
	max := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

// dati x numeri restituise il minore
func Min(a ...int) int {
	min := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > min {
			min = a[i]
		}
	}
	return min
}

// data una stringa retituisce se è lowercase
func IsLower(s string) bool {
	low := strings.ToLower(s)
	return low == s
}
