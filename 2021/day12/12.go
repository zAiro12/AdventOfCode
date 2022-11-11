package main

import "fmt"

type node struct {
	val  string
	next *node
}
type Lista struct {
	testa *node
}

func main() {
	var in string
	var albero Lista
	for {
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		aggiungi(albero, in)
	}
	stampa(albero)
}

func newNodo(val string) *node {
	return &node{val, nil}
}

func aggiungi(l Lista, val string) Lista {
	nodo := newNodo(val)
	nodo.next = l.testa
	l.testa = nodo
	return l
}

func stampa(albero Lista) {
	var next *node
	for next != nil {
		fmt.Print(next.val, " ")
		next = next.next
	}
	fmt.Println()
}
