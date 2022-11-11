package main

import (
	"fmt"
	"strconv"
	"strings"
)

type numeri struct {
	val int
	bul bool
}

var matrice [][]numeri
var c int

func main() {
	for {
		var in string
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		sub := strings.Split(in, "")
		var fin []numeri
		for i := 0; i < len(sub); i++ {
			var r numeri
			r.val, _ = strconv.Atoi(sub[i])
			r.bul = false
			fin = append(fin, r)
		}
		matrice = append(matrice, fin)
	}
	for step := 0; step < 100; step++ {
		inizio() //azzero i bool e incremento di 1
		for i := 0; i < 100; i++ {
			flash()
		}
		fine()
	}
	fmt.Println(c)
}

func incrementa(i, j int) {
	switch i {
	case 0: //CASO A B C
		switch j {
		//A
		case 0:
			matrice[i+1][j].val++
			matrice[i][j+1].val++
			matrice[i+1][j+1].val++

		//C
		case len(matrice[i]) - 1:
			matrice[i+1][j].val++
			matrice[i][j-1].val++
			matrice[i+1][j-1].val++

		//B
		default:
			matrice[i][j-1].val++
			matrice[i+1][j-1].val++
			matrice[i+1][j].val++
			matrice[i+1][j+1].val++
			matrice[i][j+1].val++
		}
	case len(matrice) - 1: //caso G H L
		switch j {
		//G
		case 0:
			matrice[i-1][j].val++
			matrice[i][j+1].val++
			matrice[i-1][j+1].val++

		//L
		case len(matrice[i]) - 1:
			matrice[i-1][j].val++
			matrice[i][j-1].val++
			matrice[i-1][j-1].val++

		//H
		default:
			matrice[i][j-1].val++
			matrice[i-1][j-1].val++
			matrice[i-1][j].val++
			matrice[i-1][j+1].val++
			matrice[i][j+1].val++
		}

	default: //caso D E F
		switch j {
		//D
		case 0:
			matrice[i-1][j].val++
			matrice[i-1][j+1].val++
			matrice[i][j+1].val++
			matrice[i+1][j+1].val++
			matrice[i+1][j].val++

		//F
		case len(matrice[i]) - 1:
			matrice[i-1][j].val++
			matrice[i-1][j-1].val++
			matrice[i][j-1].val++
			matrice[i+1][j-1].val++
			matrice[i+1][j].val++

		//E
		default:
			matrice[i][j-1].val++
			matrice[i-1][j-1].val++
			matrice[i-1][j].val++
			matrice[i-1][j+1].val++
			matrice[i][j+1].val++
			matrice[i+1][j+1].val++
			matrice[i+1][j].val++
			matrice[i+1][j-1].val++
		}
	}
}

func inizio() {
	for i := 0; i < len(matrice); i++ {
		for j := 0; j < len(matrice[i]); j++ {
			matrice[i][j].bul = false
			matrice[i][j].val++
		}
	}
}

func flash() {
	for i := 0; i < len(matrice); i++ {
		for j := 0; j < len(matrice[i]); j++ {
			if matrice[i][j].val > 9 && matrice[i][j].bul == false {
				incrementa(i, j)
				matrice[i][j].bul = true
				c++
			}
		}
	}
}

func fine() {
	for i := 0; i < len(matrice); i++ {
		for j := 0; j < len(matrice[i]); j++ {
			if matrice[i][j].val > 9 {
				matrice[i][j].val = 0
			}
		}
	}
}
