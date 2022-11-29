package utile

import "fmt"

type NodoAlbero struct {
	Destra   *NodoAlbero
	Sinistra *NodoAlbero
	Val      int
}

type Albero struct {
	Root *NodoAlbero
}

// Alloca nuovo spazio per una nuova foglia dell'albero
func NewFoglia(val int) *NodoAlbero {
	return &NodoAlbero{nil, nil, val}
}

// dato un array genera un albero completo o semicompleto
func GenAlbero(arr []int, i int) (root *NodoAlbero) {
	if i >= len(arr) {
		return nil
	}
	root = NewFoglia(arr[i])
	root.Sinistra = GenAlbero(arr, 2*i+1)
	root.Destra = GenAlbero(arr, 2*i+2)
	return root
}

//stampa un albero in preorder
func StanmpaPreorder(node *NodoAlbero) {
	//deepSearch
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	StanmpaPreorder(node.Sinistra)
	StanmpaPreorder(node.Destra)

}

//stampa un albero in inorder
func StampaInorder(node *NodoAlbero) {
	//figlio-padre-figlio
	if node == nil {
		return
	}
	StampaInorder(node.Sinistra)
	fmt.Println(node.Val, " ")
	StampaInorder(node.Destra)
}

//stampa un albero in postorder
func StampaPostorder(node *NodoAlbero) {
	//figli, radice
	if node == nil {
		return
	}
	StampaPostorder(node.Sinistra)
	StampaPostorder(node.Destra)
	fmt.Println(node.Val, " ")
}

//stampa un albero a sommario
func StampaAlberoASommario(node *NodoAlbero, spazi int) {
	for i := 0; i < spazi; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	fmt.Println(node.Val)

	if node.Sinistra != nil {
		StampaAlberoASommario(node.Sinistra, spazi+1)
	} else if node.Destra != nil {
		for i := 0; i < spazi+1; i++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}
	if node.Destra != nil {
		StampaAlberoASommario(node.Destra, spazi+1)
	} else if node.Sinistra != nil {
		for i := 0; i < spazi+1; i++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}
}

//stampa un albero
func StampaAlbero(node *NodoAlbero) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, "")
	if node.Sinistra == nil && node.Destra == nil {
		return
	}
	fmt.Print(" [")
	if node.Sinistra != nil {
		StampaAlbero(node.Sinistra)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.Destra != nil {
		StampaAlbero(node.Destra)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}
