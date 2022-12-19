package utile

import "fmt"

// struct con coordinate x, y
type Point struct {
	X, Y int
}

// funzione per stampa di debug
func Log(s string, a ...any) {
	fmt.Print("DEBUG: ", s, " ")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

// funzione per stampa di debug
func Logln(s string, a ...any) {
	fmt.Print("DEBUG: ", s, " ")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
	fmt.Println()
}

// funzione per stampare risultato della parte A
func StampaA(a ...any) {
	fmt.Print("A: ")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	//"Vaffanculo Luca, sei una testa di culo"
	fmt.Println()
}

// funzione per stampare risultato della parte B
func StampaB(a ...any) {
	fmt.Print("B: ")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

// funzione per stampare una mappa con chiave Point e qualsiasi valore come valore
func StampaMappa[T any](mappa map[Point]T, x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			fmt.Print(mappa[Point{j, i}])
		}
		fmt.Println()
	}
}
