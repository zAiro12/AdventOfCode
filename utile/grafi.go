package utile

import (
	"fmt"
)

type nodopila struct {
	val  int
	next *nodopila
}

type Grafo struct {
	N         int
	Adiacenti []*nodopila
}

func NuovoGrafo(n int) *Grafo {
	return &Grafo{n, make([]*nodopila, n)}
}

func LeggiGrafo(grafo *Grafo) {
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
}

func aggiungiAdiacente(nodo *nodopila, aggiungere *nodopila) {
	aggiungere.next = nodo.next
	nodo.next = aggiungere
}

func newNodo(val int) *nodopila {
	return &nodopila{val, nil}
}
