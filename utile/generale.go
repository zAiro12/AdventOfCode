package utile

import "strings"

type Point struct {
	X, Y int
}

func RimuoviElementoInt(s []int, index int) []int {
	a := s[:index]
	a = append(a, s[index+1:]...)
	return a
}
func RimuoviElementoString(s []string, index int) []string {
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

func IsLower(s string) bool {
	low := strings.ToLower(s)
	return low == s
}
