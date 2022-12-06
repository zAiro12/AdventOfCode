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

// data uan stringa restituisce se le lettere sono tutte diverse
func IsDiverse(s string) bool {
	mappa := make(map[byte]bool)
	
	for i := 0; i < len(s); i++ {
		mappa[s[i]] = true
	}

	return len(mappa) == len(s)
}
