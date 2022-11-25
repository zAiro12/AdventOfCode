package utile

import (
	"fmt"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type Grafo struct {
	N         int
	Adiacenti []*zairo.Testa
}

func NuovoGrafo(n int) *Grafo {
	return &Grafo{8, nil}
}

func LeggiGrafo() *Grafo {
	var num int
	fmt.Scan(&num)
	grafo := NuovoGrafo(num)
	grafo.Adiacenti = make([]*zairo.Testa, num)
	fmt.Println(len(grafo.Adiacenti))
	return grafo
}
