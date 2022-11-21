package main

import (
	"fmt"
	"github.com/zAiro12/dir"
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
}

func readTree(s string) tNode {
	return tNode{}
}

func printTree(nodo *node) string {
	var s string

	return s
}
