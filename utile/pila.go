package utile

import "fmt"

type node struct {
	val  int
	next *node
}

type testa struct {
	primo *node
}

func NewNode(val int) *node {
	return &node{val, nil}
}

func Push(val int, pila *testa) {
	nodo := NewNode(val)
	nodo.next = pila.primo
	pila.primo = nodo
}

func Pop(pila *testa) *node {
	nodo := pila.primo
	pila.primo = pila.primo.next
	return nodo
}

func (pila testa) Stampa() {
	appoggio := pila.primo
	for appoggio.next != nil {
		fmt.Print(appoggio.val, " ")
		appoggio = appoggio.next
	}
	fmt.Println()
}
