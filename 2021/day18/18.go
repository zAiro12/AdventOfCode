package main

import (
	"fmt"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type node struct {
	left  *node
	right *node
	val   int
}

type tNode struct {
	testa *node
}

func main() {
	albero := readTree("[[[1,2],3]")
	fmt.Println(printTree(albero.testa))
	var s zairo.nodo
}

func readTree(s string) tNode {
	return tNode{}
}

func printTree(nodo *node) string {
	var s string

	return s
}
