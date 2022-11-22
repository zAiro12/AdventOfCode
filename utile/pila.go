package utile

import "fmt"

type NodoPila struct {
	val  int
	next *NodoPila
}

type Testa struct {
	primo *NodoPila
}

func NewNodo(val int) *NodoPila {
	return &NodoPila{val, nil}
}

func Push(val int, pila *Testa) {
	nodo := NewNodo(val)
	nodo.next = pila.primo
	pila.primo = nodo
}

func Pop(pila *Testa) *NodoPila {
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
