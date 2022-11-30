package utile

import "strings"

// data una stringa retituisce se è lowercase
func IsLower(s string) bool {
	low := strings.ToLower(s)
	return low == s
}

// data una stringa retituisce se è uppercase
func IsUpper(s string) bool {
	up := strings.ToUpper(s)
	return up == s
}

// data una slice di qualsiasi tipo elimina l'elemento in posizione index
func RimuoviElemento[T comparable](arr []T, index int) []T {
	a := arr[:index]
	a = append(a, arr[index+1:]...)
	return a
}
