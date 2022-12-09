package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type albero struct {
	val    int
	siVede bool
}

func main() {
	var matrice [][]albero

	input := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(input)

	var in string
	var split []string
	var num int

	for sc.Scan() {
		in = sc.Text()

		split = strings.Split(in, "")

		var riga []albero
		for i := 0; i < len(split); i++ {
			num, _ = strconv.Atoi(split[i])
			riga = append(riga, albero{num, false})
		}
		matrice = append(matrice, riga)
	}

	zairo.StampaA(parteA(matrice))
	zairo.StampaB(parteB(matrice))
}

func stampaMatrice(matrice [][]albero) {
	for i := 0; i < len(matrice); i++ {
		for j := 0; j < len(matrice[i]); j++ {
			fmt.Print(matrice[i][j])
		}
		fmt.Println()
	}
}

func parteA(matrice [][]albero) int {
	for i := 0; i < len(matrice); i++ {
		for j := 0; j < len(matrice[i]); j++ {
			cerca(matrice, i, j)
		}
	}

	stampaMatrice(matrice)

	return cercaT(matrice)
}

func parteB(matrice [][]albero) int {
	var max int
	for i := 0; i < len(matrice); i++ {
		for j := 0; j < len(matrice[i]); j++ {
			temp := guarda(matrice, i, j)
            if temp > max{
                max = temp
            }
		}
	}
	return max
}

func cerca(matrice [][]albero, x, y int) {
	if x == 0 || x == len(matrice)-1 || y == 0 || y == len(matrice)-1 {
		matrice[x][y].siVede = true
	} else if cercaDestra(matrice, x, y) || cercaSu(matrice, x, y) || cercaSinistra(matrice, x, y) || cercaGiu(matrice, x, y) {
		matrice[x][y].siVede = true
	}
}

func cercaDestra(matrice [][]albero, riga, colonna int) bool {
	for i := colonna + 1; i < len(matrice[riga]); i++ {
		if matrice[riga][colonna].val <= matrice[riga][i].val {
			return false
		}
	}
	return true
}

func cercaSu(matrice [][]albero, riga, colonna int) bool {
	for i := riga - 1; i >= 0; i-- {
		if matrice[riga][colonna].val <= matrice[i][colonna].val {
			return false
		}
	}
	return true
}

func cercaSinistra(matrice [][]albero, riga, colonna int) bool {
	for i := colonna - 1; i >= 0; i-- {
		if matrice[riga][colonna].val <= matrice[riga][i].val {
			return false
		}
	}
	return true
}

func cercaGiu(matrice [][]albero, riga, colonna int) bool {
	for i := riga + 1; i < len(matrice); i++ {
		if matrice[riga][colonna].val <= matrice[i][colonna].val {
			return false
		}
	}
	return true
}

func cercaT(matrice [][]albero) (c int) {
	for i := 0; i < len(matrice); i++ {
		for j := 0; j < len(matrice[i]); j++ {
			if matrice[i][j].siVede {
				c++
			}
		}
	}
	return
}

func guarda(matrice [][]albero, riga, colonna int) int {
	var lati [4]int

	//destra
	for i := colonna + 1; i < len(matrice[riga]); i++ {
		lati[0]++
		if matrice[riga][colonna].val <= matrice[riga][i].val {
			break
		}
	}

	//su
	for i := riga - 1; i >= 0; i-- {
		lati[1]++
		if matrice[riga][colonna].val <= matrice[i][colonna].val {
			break
		}
	}

	//sinistra
	for i := colonna - 1; i >= 0; i-- {
		lati[2]++
		if matrice[riga][colonna].val <= matrice[riga][i].val {
			break
		}
	}

	//giu
	for i := riga + 1; i < len(matrice); i++ {
		lati[3]++
		if matrice[riga][colonna].val <= matrice[i][colonna].val {
			break
		}
	}
	return zairo.MoltipicaSlice(lati[0:4])
}
