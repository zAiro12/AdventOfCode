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
			matrice[zairo.Point{X: x, Y: y}] = dato{in[x], 99}
		}
		y++
	}

	//StampaMappa(matrice, x, y)
	zairo.StampaA(parteA(cercaS(matrice), matrice))
	zairo.StampaB(parteB(matrice))
}

func StampaMappa(mappa map[zairo.Point]dato, x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			fmt.Printf("%c", mappa[zairo.Point{X: j, Y: i}].nome)
		}
		fmt.Println()
	}
}
func StampaMappaVal(mappa map[zairo.Point]dato, x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			fmt.Printf("%c %2d ", mappa[zairo.Point{X: j, Y: i}].nome, mappa[zairo.Point{X: j, Y: i}].val)
		}
		fmt.Println()
	}
	fmt.Println()
}

func cercaS(matrice map[zairo.Point]dato) zairo.Point {
	for k, v := range matrice {
		if v.nome == 'S' {
			return k
		}
	}
	return zairo.Point{}
}

func parteA(start zairo.Point, matrice map[zairo.Point]dato) int {
	var frontiera []zairo.Point
	exl := make(map[zairo.Point]bool)

	frontiera = append(frontiera, start)

	var daAnalizzare zairo.Point
	for len(frontiera) != 0 {
		daAnalizzare = frontiera[0]
		frontiera = frontiera[1:]

		if matrice[daAnalizzare].nome == 'S' {
			matrice[daAnalizzare] = dato{'a', 0}
		}

		if !exl[daAnalizzare] {

			if matrice[daAnalizzare].nome == 'E' {
				return matrice[daAnalizzare].val
			}

			exl[daAnalizzare] = true

			// UP
			if ((matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y - 1}].nome == 'E' && matrice[daAnalizzare].nome == 'z') || (matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y - 1}].nome <= matrice[daAnalizzare].nome+1 && matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y - 1}].nome != 0 && matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y - 1}].nome != 'E')) && !exl[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y - 1}] {

				frontiera = append(frontiera, zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y - 1})
				matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y - 1}] = dato{matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y - 1}].nome, matrice[daAnalizzare].val + 1}
			}

			// DOWN
			if ((matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y + 1}].nome == 'E' && matrice[daAnalizzare].nome == 'z') || (matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y + 1}].nome <= matrice[daAnalizzare].nome+1 && matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y + 1}].nome != 0 && matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y + 1}].nome != 'E')) && !exl[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y + 1}] {

				frontiera = append(frontiera, zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y + 1})
				matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y + 1}] = dato{matrice[zairo.Point{X: daAnalizzare.X, Y: daAnalizzare.Y + 1}].nome, matrice[daAnalizzare].val + 1}
			}

			// LEFT
			if ((matrice[zairo.Point{X: daAnalizzare.X - 1, Y: daAnalizzare.Y}].nome == 'E' && matrice[daAnalizzare].nome == 'z') || (matrice[zairo.Point{X: daAnalizzare.X - 1, Y: daAnalizzare.Y}].nome <= matrice[daAnalizzare].nome+1 && matrice[zairo.Point{X: daAnalizzare.X - 1, Y: daAnalizzare.Y}].nome != 0 && matrice[zairo.Point{X: daAnalizzare.X - 1, Y: daAnalizzare.Y}].nome != 'E')) && !exl[zairo.Point{X: daAnalizzare.X - 1, Y: daAnalizzare.Y}] {

				frontiera = append(frontiera, zairo.Point{X: daAnalizzare.X - 1, Y: daAnalizzare.Y})
				matrice[zairo.Point{X: daAnalizzare.X - 1, Y: daAnalizzare.Y}] = dato{matrice[zairo.Point{X: daAnalizzare.X - 1, Y: daAnalizzare.Y}].nome, matrice[daAnalizzare].val + 1}
			}

			// RIGHT
			if ((matrice[zairo.Point{X: daAnalizzare.X + 1, Y: daAnalizzare.Y}].nome == 'E' && matrice[daAnalizzare].nome == 'z') || (matrice[zairo.Point{X: daAnalizzare.X + 1, Y: daAnalizzare.Y}].nome <= matrice[daAnalizzare].nome+1 && matrice[zairo.Point{X: daAnalizzare.X + 1, Y: daAnalizzare.Y}].nome != 0 && matrice[zairo.Point{X: daAnalizzare.X + 1, Y: daAnalizzare.Y}].nome != 'E')) && !exl[zairo.Point{X: daAnalizzare.X + 1, Y: daAnalizzare.Y}] {

				frontiera = append(frontiera, zairo.Point{X: daAnalizzare.X + 1, Y: daAnalizzare.Y})
				matrice[zairo.Point{X: daAnalizzare.X + 1, Y: daAnalizzare.Y}] = dato{matrice[zairo.Point{X: daAnalizzare.X + 1, Y: daAnalizzare.Y}].nome, matrice[daAnalizzare].val + 1}
			}

		}
	}

	return 0
}

func parteB(matrice map[zairo.Point]dato) int {
	var min int = math.MaxInt
	var tmp int

	for k, v := range matrice {
		for k1, v1 := range matrice {
			v1 = dato{v1.nome, 0}
			matrice[k1] = v1
		}
		if v.nome == 'a' || v.nome == 'S' {
			tmp = parteA(k, matrice)

			if tmp < min && tmp != 0 {
				min = tmp
			}
		}
	}
	return min
}
