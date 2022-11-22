package utile

import "fmt"

type Node struct {
	val  int
	next *Node
}

type Testa struct {
	primo *Node
}

func NewNode(val int) *Node {
	return &Node{val, nil}
}

func Push(val int, pila *Testa) {
	nodo := NewNode(val)
	nodo.next = pila.primo
	pila.primo = nodo
}

func Pop(pila *Testa) *Node {
	nodo := pila.primo
	pila.primo = pila.primo.next
	return nodo
}

func (pila Testa) Stampa() {
	appoggio := pila.primo
	for appoggio.next != nil {
		fmt.Print(appoggio.val, " ")
		appoggio = appoggio.next
	}
	fmt.Println()
}
