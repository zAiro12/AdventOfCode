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
	nRighe := len(matrice)
	nColonne := len(matrice[0])
	res := bestToGo(matrice)
	fmt.Println("MATRICE ORIGINALE")
	stampa(matrice)
	fmt.Println()
	fmt.Println("MATRICE DELLE SOMME")
	stampa(res)
	fmt.Println()
	fmt.Println("OTTIMO:", res[nRighe - 1][nColonne - 1])
	fmt.Println("PERCORSO:", getSnake(matrice, res, len(matrice) - 1, len(matrice[0]) - 1))
}

func min(a,b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func bestToGo(matrice [][]int) (res [][]int) {
	nRighe := len(matrice)
	nColonne := len(matrice[0])
	// Creo matrice risultato
	res = make([][]int, nRighe)
	for i := 0; i < nRighe; i++ {
		res[i] = make([]int, nColonne)
	}
	// Primo elemento (in alto a sx)
	res[0][0] = matrice[0][0]
	// Prima riga
	for j := 1; j < nColonne; j++ {
		res[0][j] = res[0][j-1] + matrice[0][j] 
	}
	// Prima colonna
	for i := 1; i < nRighe; i++ {
		res[i][0] = res[i-1][0] + matrice[i][0] 
	}
	// Altri elementi
	for i := 1; i < nRighe; i++ {
		for j := 1; j < nColonne; j++ {
			res[i][j] = min(res[i-1][j], res[i][j-1]) + matrice[i][j]
		}
	}
	return
}

func getSnake(matrice [][]int, res [][]int, i, j int) (m string) {
	nRighe := len(matrice)
	nColonne := len(matrice[0])
	i = 0
	j = 0	
	for i < nRighe - 1 && j < nColonne - 1 {
		if res[i][j] +  matrice[i][j] == res[i+1][j] {
			m += "D"
			i++
		} else {
			m += "R"
			j++
		}
	}
	for i < nRighe - 1 {
		m += "D"
		i++
	} 
	for j < nColonne - 1 {
		m += "R"
		j++
	} 
	return
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
