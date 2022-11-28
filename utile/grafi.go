package utile

import (
	"fmt"
	"strings"
)

type NodopilaInt struct {
	Val  int
	Next *NodopilaInt
}

type GrafoInt struct {
	Len       int
	Adiacenti []*NodopilaInt
}

type GrafoString map[string][]string

//alloca memoria per un nuovo grafo vuoto
func NuovoGrafoInt(n int) *GrafoInt {
	return &GrafoInt{n, make([]*NodopilaInt, n)}
}

//aggiunge ad un grafo vuoto nuove coppie da standart input
func AggiungiCoppieGrafoInt(grafo *GrafoInt) {
	var x, y int
	for {
		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}
		fmt.Scan(&y)
		inputCoppieInt(grafo, x, y)
	}
}

//aggiunge ad un grafo una coppia
func inputCoppieInt(grafo *GrafoInt, x, y int) {
	nodo := newNodoInt(y)
	if grafo.Adiacenti[x] != nil {
		aggiungiAdiacenteInt(grafo.Adiacenti[x], nodo)
	} else {
		grafo.Adiacenti[x] = nodo
	}
	// fmt.Println("debug", x, y)
	// fmt.Println("DEBUG:", grafo.Adiacenti)
}

//crea un arco con un nodo già esistente
func aggiungiAdiacenteInt(nodo *NodopilaInt, aggiungere *NodopilaInt) {
	aggiungere.Next = nodo.Next
	nodo.Next = aggiungere
}

//alloca spazio per un vuono nodo
func newNodoInt(val int) *NodopilaInt {
	return &NodopilaInt{val, nil}
}

//stampa il grafo nodo per nodo
func StampaGrafoInt(grafo *GrafoInt) {
	var appoggio *NodopilaInt
	for i := 0; i < grafo.Len; i++ {
		fmt.Println("VERTICE", i)
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

//restituisce se due nodi sono collegati da un arco o no
func IsArcoGrafoInt(grafo *GrafoInt, x, y int) bool {
	if x > grafo.Len || x < 0 {
		return false
	}
	appoggio := grafo.Adiacenti[x]
	for appoggio != nil {
		if appoggio.Val == y {
			return true
		}
		appoggio = appoggio.Next
	}
	return false
}

//alloca spazio per un nuovo grafo vuoto
func NewGrafoString() GrafoString {
	return make(GrafoString, 0)
}

//metodo per inserire da standard input nuovi nodi al grafo
func (grafo GrafoString) InputString() {
	var in string
	for {
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		split := strings.Split(in, "-")
		grafo[split[0]] = append(grafo[split[0]], split[1])
		grafo[split[1]] = append(grafo[split[1]], split[0])
	}
}

//metodo per stapare il grafo
func (grafo GrafoString) StampaGrafoString() {
	for k, v := range grafo {
		fmt.Println("VERTICE:", k)
		for i := 0; i < len(v); i++ {
			fmt.Print(v[i], " ")
		}
		fmt.Println()
	}
}

//metodo per stapare gli archi collegati ad un nodo
func (grafo GrafoString) Archi(v string) []string {
	return grafo[v]
}
