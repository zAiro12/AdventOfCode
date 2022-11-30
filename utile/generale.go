package utile

import "fmt"

// struct con coordinate x, y
type Point struct {
	X, Y int
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
func RimuoviElemento(arr []any, index int) []any {
	a := arr[:index]
	a = append(a, arr[index+1:]...)
	return a
}

// funzione per stampa di debug
func Log(a ...any) {
	fmt.Print("DEBUG: ")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

//funzione per stampare risultato della parte A
func StampaA(a ...any){
	fmt.Print("A: ")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

//funzione per stampare risultato della parte B
func StampaB(a ...any){
	fmt.Print("A: ")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}