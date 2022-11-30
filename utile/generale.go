package utile

import "fmt"

// struct con coordinate x, y
type Point struct {
	X, Y int
}

// funzione per stampa di debug
func Log(a ...any) {
	fmt.Print("DEBUG: ")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

// funzione per stampare risultato della parte A
func StampaA(a ...any) {
	fmt.Print("A: ")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

// funzione per stampare risultato della parte B
func StampaB(a ...any) {
	fmt.Print("A: ")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}
