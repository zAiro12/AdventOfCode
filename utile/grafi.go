package utile

import (
	"fmt"
)

type Nodopila struct {
	Val  int
	Next *Nodopila
}

type Grafo struct {
	Len       int
	Adiacenti []*Nodopila
}

func NuovoGrafo(n int) *Grafo {
	return &Grafo{n, make([]*Nodopila, n)}
}

func AggiungiCoppieGrafo(grafo *Grafo) {
	var x, y int
	for {
		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}
		fmt.Scan(&y)
		inputCoppie(grafo, x, y)
	}
}

func inputCoppie(grafo *Grafo, x, y int) {
	nodo := newNodo(y)
	if grafo.Adiacenti[x] != nil {
		aggiungiAdiacente(grafo.Adiacenti[x], nodo)
	} else {
		grafo.Adiacenti[x] = nodo
	}
	fmt.Println("debug", x, y)
	fmt.Println("DEBUG:", grafo.Adiacenti)
}

func aggiungiAdiacente(nodo *Nodopila, aggiungere *Nodopila) {
	aggiungere.Next = nodo.Next
	nodo.Next = aggiungere
}

func newNodo(val int) *Nodopila {
	return &Nodopila{val, nil}
}

func StampaGrafo(grafo *Grafo) {
	var appoggio *Nodopila
	for i := 0; i < grafo.Len; i++ {
		fmt.Println("NODO", i)
		if grafo.Adiacenti[i] != nil {
			appoggio = grafo.Adiacenti[i]
			fmt.Print(appoggio.Val, " ")
			for appoggio.Next != nil {
				appoggio = appoggio.Next
				fmt.Print(appoggio.Val, " ")
			}
		}
		fmt.Println()
	}
}
