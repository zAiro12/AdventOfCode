package utile

import (
	"fmt"
)

type Lista struct {
	Val  int
	Next *Lista
}
type Grafo struct {
	N         int
	Adiacenti []*Lista
}

func NuovoGrafo(n int) *Grafo {
	return &Grafo{8, nil}
}

func LeggiGrafo() *Grafo {
	var num int
	fmt.Scan(&num)
	grafo := NuovoGrafo(num)
	grafo.Adiacenti = make([]*Lista, num)
	fmt.Println(len(grafo.Adiacenti))
	return grafo
}
