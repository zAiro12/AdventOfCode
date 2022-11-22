package main

import (
	"fmt"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	albero := ReadTree("[[[1,2],3]")
	fmt.Println(PrintTree(albero.Root))
	fmt.Println(zairo.Max(1, 2))
}

func ReadTree(s string) Albero {
	return Albero{nil}
}

func PrintTree(nodo *NodoAlbero) string {
	var s string

	return s
}
