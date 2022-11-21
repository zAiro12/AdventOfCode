package utile

import "fmt"

type node struct {
	val  int
	next *node
}

type testa struct {
	primo *node
}

func newNode(val int) *node {
	return &node{val, nil}
}

func aggiungi(val int, pila *testa) {
	nodo := newNode(val)
	nodo.next = pila.primo
	pila.primo = nodo
}

func pop(pila *testa) *node {
	nodo := pila.primo
	pila.primo = pila.primo.next
	return nodo
}

func (pila testa) stampa() {
	appoggio := pila.primo
	for appoggio.next != nil {
		fmt.Print(appoggio.val, " ")
		appoggio = appoggio.next
	}
	fmt.Println()
}
