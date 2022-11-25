package utile

import (
	"fmt"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type Grafo struct {
	N         int
	Adiacenti []*zairo.NodoPila
}

func NuovoGrafo(n int) *Grafo {
	return &Grafo{8, nil}
}

func LeggiGrafo() *Grafo {
	var num int
	fmt.Scan(&num)
	grafo := NuovoGrafo(num)
	grafo.Adiacenti = make([]*zairo.NodoPila, num)
	var x, y int
	for {
		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}
		fmt.Scan(&y)
		inputCoppie(grafo, x, y)
	}
	return grafo
}

func inputCoppie(grafo *Grafo, x, y int) {
	if grafo.Adiacenti[x] != nil {
		aggiungiAdiacente(x, y)
	} else {
		grafo.Adiacenti[x] = zairo.NewNodo(y)
	}
}

func aggiungiAdiacente(grafo *Grafo, x, y int) {
	appoggio := grafo.Adiacenti[x].Next
	for appoggio.Next != nil {
		appoggio = appoggio.Next
	}
	nuovo := zairo.NewNodoPila()
	appoggio.Next = nuovo
}
