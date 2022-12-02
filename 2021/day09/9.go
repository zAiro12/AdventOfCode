package main

import (
	"fmt"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type punto struct {
	x, y int
}
type risultato struct {
	val  int
	flag bool
}

func main() {
	matrice := make(map[punto]risultato)
	var in string
	var appoggio punto
	var lunghezza int
	for i := 0; ; i++ {
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		appoggio.x = i
		lunghezza = i + 1
		for j := 0; j < len(in); j++ {
			appoggio.y = j
			matrice[appoggio] = risultato{int(in[j] - '0'), false}
		}
	}
	zairo.Log("L:", lunghezza)
	// stampa(matrice, lunghezza)
	zairo.StampaA(parteA(matrice, 10))
}

func parteA(matrice map[punto]risultato, lunghezza int) (c int) {
	lunghezzaY := 5
	lunghezzaX := 10

	for y := 0; y < lunghezzaY; y++ {
		for x := 0; x < lunghezzaX; x++ {

			//zairo.Log("x,y: ", x, y)

			if x == 0 {

				if y == 0 {
					if matrice[punto{x, y}].val < matrice[punto{x + 1, y}].val && matrice[punto{x, y}].val < matrice[punto{x, y + 1}].val {
						c += 1 + matrice[punto{x, y}].val
						zairo.Log(matrice[punto{x, y}].val)
					}
				} else if y == lunghezzaY-1 {
					if matrice[punto{x, y}].val < matrice[punto{x + 1, y}].val && matrice[punto{x, y}].val < matrice[punto{x, y - 1}].val {
						c += 1 + matrice[punto{x, y}].val
						zairo.Log(matrice[punto{x, y}].val)
					}
				} else {
					if matrice[punto{x, y}].val < matrice[punto{x + 1, y}].val && matrice[punto{x, y}].val < matrice[punto{x, y + 1}].val && matrice[punto{x, y}].val < matrice[punto{x, y - 1}].val {
						c += 1 + matrice[punto{x, y}].val
						zairo.Log(matrice[punto{x, y}].val)
					}
				}

			} else if x == lunghezzaX-1 {

				if y == 0 {
					if matrice[punto{x, y}].val < matrice[punto{x - 1, y}].val && matrice[punto{x, y}].val < matrice[punto{x, y + 1}].val {
						c += 1 + matrice[punto{x, y}].val
						zairo.Log(matrice[punto{x, y}].val)
					}
				} else if y == lunghezzaY-1 {
					if matrice[punto{x, y}].val < matrice[punto{x - 1, y}].val && matrice[punto{x, y}].val < matrice[punto{x, y - 1}].val {
						c += 1 + matrice[punto{x, y}].val
						zairo.Log(matrice[punto{x, y}].val)
					}
				} else {
					if matrice[punto{x, y}].val < matrice[punto{x - 1, y}].val && matrice[punto{x, y}].val < matrice[punto{x, y + 1}].val && matrice[punto{x, y}].val < matrice[punto{x, y - 1}].val {
						c += 1 + matrice[punto{x, y}].val
						zairo.Log(matrice[punto{x, y}].val)
					}
				}

			} else {

				if y == 0 {
					if matrice[punto{x, y}].val < matrice[punto{x + 1, y}].val && matrice[punto{x, y}].val < matrice[punto{x, y + 1}].val && matrice[punto{x, y}].val < matrice[punto{x - 1, y}].val {
						c += 1 + matrice[punto{x, y}].val
						zairo.Log(matrice[punto{x, y}].val)
					}
				} else if y == lunghezzaY-1 {
					if matrice[punto{x, y}].val < matrice[punto{x + 1, y}].val && matrice[punto{x, y}].val < matrice[punto{x, y - 1}].val && matrice[punto{x, y}].val < matrice[punto{x - 1, y}].val {
						c += 1 + matrice[punto{x, y}].val
						zairo.Log(matrice[punto{x, y}].val)
					}
				} else {
					if matrice[punto{x, y}].val < matrice[punto{x + 1, y}].val && matrice[punto{x, y}].val < matrice[punto{x, y + 1}].val && matrice[punto{x, y}].val < matrice[punto{x, y - 1}].val && matrice[punto{x, y}].val < matrice[punto{x - 1, y}].val {
						c += 1 + matrice[punto{x, y}].val
						zairo.Log(matrice[punto{x, y}].val)
					}
				}

			}
		}
	}
	return
}

func area(matrice map[punto]risultato, x, y int) int {
	if matrice[punto{x, y}].flag || matrice[punto{x, y}].val == 9 || matrice[punto{x, y}].val == 0 {
		return 0
	}
	matrice[punto{x, y}] = risultato{matrice[punto{x, y}].val, true}
	return 1 + area(matrice, x+1, y) + area(matrice, x, y+1) + area(matrice, x-1, y) + area(matrice, x, y-1)
}

func stampa(matrice map[punto]risultato, lunghezza int) {
	for i := 0; i < lunghezza; i++ {
		for j := 0; j < lunghezza; j++ {
			fmt.Print(matrice[punto{i, j}].val)
		}
		fmt.Println()
	}
}
