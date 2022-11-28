package main

import (
	"fmt"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	fmt.Println("ciao")

	var a, b int
	a = 1
	b = 5

	c := a + b
	fmt.Println(c)
	_ = zairo.Max(a, b, c)
}
