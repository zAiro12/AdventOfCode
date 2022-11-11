package main

import (
	"fmt"
	"strconv"
	"strings"
)

var c int

func main() {
	var matrice [][]int
	for {
		var in string
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		r := strings.Split(in, "")
		var in2 []int
		for i := 0; i < len(r); i++ {
			var in3 int
			in3, _ = strconv.Atoi(r[i])
			in2 = append(in2, in3)
		}
		matrice = append(matrice, in2)
	}
	k := matrice[0][0]
	fmt.Println(start(matrice) + k)

}

//funzione ricorsiva che crea una sub slice
func start(matrice [][]int) int {
	if len(matrice) == 1 {
		return c
	}
	if len(matrice) == 2 && len(matrice[0]) == 2 {
		return c
	}
	if matrice[0][1] > matrice[1][0] {
		c += matrice[0][1]
		matrice = riga(matrice)
	} else {
		c += matrice[1][0]
		matrice = colonna(matrice)
	}
	stampa(matrice)
	fmt.Println()
	return c + start(matrice)
}

//elimina riga
func riga(matrice [][]int) [][]int {
	return matrice[1:]
}

//elimina colonna
func colonna(matrice [][]int) [][]int {
	var sub [][]int
	for i := 0; i < len(matrice); i++ {
		sub = append(sub, matrice[i][1:])
	}
	return sub
}

//stampa matrice a forma di matrice
func stampa(matrice [][]int) {
	for i := 0; i < len(matrice); i++ {
		for j := 0; j < len(matrice[i]); j++ {
			fmt.Print(matrice[i][j], " ")
		}
		fmt.Println()
	}
}
