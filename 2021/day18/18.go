package main

import (
	"fmt"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	albero := readTree("[[[1,2],3]")
	fmt.Println(printTree(albero.root))
}

func readTree(s string) zairo.Albero {
	return zairo.Albero{}
}

func printTree(nodo *zairo.NodoAlbero) string {
	var s string

	return s
}
