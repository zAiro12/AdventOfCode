package utile

import "fmt"

type NodoPila struct {
	Val  int
	Next *NodoPila
}

type Testa struct {
	Primo *NodoPila
}

//alloca spazio per un nuovo nodo
func NewNodoPila(val int) *NodoPila {
	return &NodoPila{val, nil}
}

//aggiunge in elemento in testa alla pila
func PushPila(val int, pila *Testa) {
	nodo := NewNodoPila(val)
	nodo.Next = pila.Primo
	pila.Primo = nodo
}

//toglie l'elemento in testa alla pila
func PopPila(pila *Testa) *NodoPila {
	nodo := pila.Primo
	pila.Primo = pila.Primo.Next
	return nodo
}

//metodo per stampare la pila
func (pila Testa) StampaPila() {
	appoggio := pila.Primo
	for appoggio.Next != nil {
		fmt.Print(appoggio.Val, " ")
		appoggio = appoggio.Next
	}
	fmt.Println()
}
