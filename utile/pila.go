package utile

import "fmt"

type NodoPila struct {
	Val  int
	Next *NodoPila
}

type Testa struct {
	Primo *NodoPila
}

func NewNodo(val int) *NodoPila {
	return &NodoPila{val, nil}
}

func Push(val int, pila *Testa) {
	nodo := NewNodo(val)
	nodo.Next = pila.Primo
	pila.Primo = nodo
}

func Pop(pila *Testa) *NodoPila {
	nodo := pila.Primo
	pila.Primo = pila.Primo.Next
	return nodo
}

func (pila Testa) Stampa() {
	appoggio := pila.Primo
	for appoggio.Next != nil {
		fmt.Print(appoggio.Val, " ")
		appoggio = appoggio.Next
	}
	fmt.Println()
}
