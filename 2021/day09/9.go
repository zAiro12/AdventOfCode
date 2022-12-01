package main

import (
	"fmt"
	"sort"
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
	var lunghezza, c int
	var finale []int
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

	//stampa(matrice)

	for i := 0; i < lunghezza; i++ {
		for j := 0; j < lunghezza; j++ {
			switch i {
			case 0:
				if j == 0 {
					if matrice[punto{i, j}].val < matrice[punto{i + 1, j}].val && matrice[punto{i, j}].val < matrice[punto{i, j + 1}].val {
						c += matrice[punto{i, j}].val
					}
				} else if j == lunghezza-1 {
					if matrice[punto{i, j}].val < matrice[punto{i + 1, j}].val && matrice[punto{i, j}].val < matrice[punto{i, j + 1}].val {
						c += matrice[punto{i, j}].val
					}
				} else {

				}
				continue

			case lunghezza - 1:
				continue

			default:
				continue

			}
		}
	}
	sort.Ints(finale)
	fmt.Println(finale)
	fmt.Println("B:", finale[len(finale)-1]*finale[len(finale)-2]*finale[len(finale)-3])
	fmt.Println(c)
}

func area(matrice map[punto]risultato, x, y int) int {
	if matrice[punto{x, y}].flag || matrice[punto{x, y}].val == 9 || matrice[punto{x, y}].val == 0 {
		return 0
	}
	matrice[punto{x, y}] = risultato{matrice[punto{x, y}].val, true}
	return 1 + area(matrice, x+1, y) + area(matrice, x, y+1) + area(matrice, x-1, y) + area(matrice, x, y-1)
}

func stampa(matrice map[punto]risultato) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			fmt.Print(matrice[punto{i, j}].val)
		}
		fmt.Println()
	}
}
