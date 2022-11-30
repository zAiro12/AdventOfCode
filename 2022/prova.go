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
	zairo.StampaA(zairo.Max(a, b, c), a, b)

	s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	j := []any{int(1), string("2"), int(3), byte('4'), float64(5.908), int(6), bool(false), int(8), int(9), 10}

	fmt.Println(zairo.RimuoviElemento(s, 4))
	fmt.Println(zairo.RimuoviElemento(i, 4))
	fmt.Println(s, i)
	fmt.Printf("%T \n", j[9])

}
