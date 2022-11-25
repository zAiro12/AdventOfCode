package utile

import (
	"fmt"
)

type nodopila struct {
	val  int
	next *nodopila
}

type Grafo struct {
	Len       int
	Adiacenti []*nodopila
}

func NuovoGrafo(n int) *Grafo {
	return &Grafo{n, make([]*nodopila, n)}
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

func aggiungiAdiacente(nodo *nodopila, aggiungere *nodopila) {
	aggiungere.next = nodo.next
	nodo.next = aggiungere
}

func newNodo(val int) *nodopila {
	return &nodopila{val, nil}
}

func StampaGrafo(grafo *Grafo) {
	var appoggio *nodopila
	for i := 0; i < grafo.Len; i++ {
		fmt.Println("NODO", i)
		if grafo.Adiacenti[i].next != nil {
			appoggio = grafo.Adiacenti[i].next
			fmt.Print(appoggio.val)
			for appoggio.next != nil {
				fmt.Println(appoggio.val)
				appoggio = appoggio.next
			}
		}
		fmt.Println()
	}
}
