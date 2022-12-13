package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type dato struct {
	nome byte
	val  int
}

func main() {
	matrice := make(map[zairo.Point]dato)

	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)

	var in string

	var y, x int

	for sc.Scan() {
		in = sc.Text()

		for x = 0; x < len(in); x++ {
			matrice[zairo.Point{X: x, Y: y}] = dato{in[x], math.MaxInt}
		}
		y++
	}

	StampaMappa(matrice, x, y)
	zairo.StampaA(cerca(cercaS(matrice), matrice, zairo.Point{X: -1, Y: -1}, 0))
	zairo.StampaB()
}

func StampaMappa(mappa map[zairo.Point]dato, x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			fmt.Printf("%c", mappa[zairo.Point{X: j, Y: i}].nome)
		}
		fmt.Println()
	}
}

func cercaS(matrice map[zairo.Point]dato) zairo.Point {
	for k, v := range matrice {
		if v.nome == 'S' {
			return k
		}
	}
	return zairo.Point{}
}

func cerca(nodo zairo.Point, matrice map[zairo.Point]dato, partenza zairo.Point, counter int) int {
	zairo.Log("ciclo", counter)

	if matrice[nodo].nome == 'E' {
		return counter
	}
	if matrice[nodo].nome == 'S' {
		matrice[nodo] = dato{'a', 0}
	}

	zairo.Log("a", counter < matrice[nodo].val)
	if counter < matrice[nodo].val || matrice[nodo].nome == 'a' {
		matrice[nodo] = dato{matrice[nodo].nome, counter}

		x := nodo.X
		y := nodo.Y

		var next zairo.Point
		//Up
		next = zairo.Point{X: x, Y: y - 1}

		if matrice[next].val != 0 && matrice[next].nome <= matrice[nodo].nome+1 && matrice[next].nome+matrice[nodo].nome == 'E'+'z' && partenza != next {
			return cerca(next, matrice, nodo, counter+1)
		}

		//Down
		next = zairo.Point{X: x, Y: y + 1}

		if matrice[next].val != 0 && matrice[next].nome <= matrice[nodo].nome+1 && matrice[next].nome+matrice[nodo].nome == 'E'+'z' && partenza != next {
			return cerca(next, matrice, nodo, counter+1)
		}

		//Left
		next = zairo.Point{X: x - 1, Y: y}

		if matrice[next].val != 0 && matrice[next].nome <= matrice[nodo].nome+1 && matrice[next].nome+matrice[nodo].nome == 'E'+'z' && partenza != next {
			return cerca(next, matrice, nodo, counter+1)
		}

		//Right
		next = zairo.Point{X: x + 1, Y: y}

		if matrice[next].val != 0 && matrice[next].nome <= matrice[nodo].nome+1 && matrice[next].nome+matrice[nodo].nome == 'E'+'z' && partenza != next {
			return cerca(next, matrice, nodo, counter+1)
		}

	}

	return math.MaxInt
}
