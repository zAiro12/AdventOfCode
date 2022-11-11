/* package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var matrice [][]bool
	var piegare []string
	var x []int
	var y []int
	for {
		var in string
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		if in == "fold" {
			fmt.Scan(&in)
			fmt.Scan(&in)
			piegare = append(piegare, in)
		} else {
			r := strings.Split(in, ",")
			s, _ := strconv.Atoi(r[0])
			x = append(x, s)
			u, _ := strconv.Atoi(r[1])
			y = append(y, u)
		}
	}
	for i := 0; i < len(x); i++ {
		switch x[i] {
		case 0:
			if y[i] == 0 {
				aggiungi(x[i], y[i], matrice)
			} else {
				aggiungi(x[i], y[i]-1, matrice)
			}
		default:
			if y[i] == 0 {
				aggiungi(x[i]-1, y[i], matrice)
			} else {
				aggiungi(x[i]-1, y[i]-1, matrice)
			}
		}
		//aggiungi(x[i], y[i])

	}
	for i := 0; i < len(piegare); i++ {
		if i == 0 {
			r := strings.Split(piegare[i], "=")
			if r[0] == "x" {
				xi, _ := strconv.Atoi(r[1])
				matrice = piegax1(xi, matrice)
			} else {
				yi, _ := strconv.Atoi(r[1])
				matrice = piegay1(yi, matrice)
			}
		} else {
			r := strings.Split(piegare[i], "=")
			if r[0] == "x" {
				xi, _ := strconv.Atoi(r[1])
				matrice = piegax(xi, matrice)
			} else {
				yi, _ := strconv.Atoi(r[1])
				matrice = piegay(yi, matrice)
			}
		}
	}

	stampanormale(matrice)
}

func aggiungi(x, y int, matrice [][]bool) [][]bool {
	matrice[x][y] = true
}

func max(slice []int) int {
	var max int
	for i := 0; i < len(slice); i++ {
		if slice[i] >= max {
			max = slice[i]
		}
	}
	return max
}

func stampa(matrice [][]bool) [][]bool {
	for i := 0; i < len(matrice); i++ {
		for j := 0; j < len(matrice[i]); j++ {
			if matrice[i][j] == false {
				fmt.Print(". ")
			} else {
				fmt.Print("# ")
			}
		}
		fmt.Println()
	}
}

func piegax1(x int, matrice [][]bool) [][]bool {
	var nuova [][]bool
	var in []bool
	for k := 0; k < x-1; k++ { //quante volte
		for j := 0; j < len(matrice[0]); j++ {
			if matrice[k][j] == true || matrice[len(matrice)-1-k][j] == true {
				in = append(in, true)
			} else {
				in = append(in, false)
			}
		}
		nuova = append(nuova, in)
	}
	return nuova
}

func piegay1(y int, matrice [][]bool) [][]bool {
	var nuova [][]bool
	for k := 0; k < y-1; k++ { //quante volte
		var in []bool
		for j := 0; j < len(matrice[0]); j++ {
			if matrice[k][j] == true || matrice[len(matrice)-1-k][j] == true {
				in = append(in, true)
			} else {
				in = append(in, false)
			}
		}
		nuova = append(nuova, in)
	}
	return nuova
}

func stampanormale(nuova [][]bool) {
	for i := 0; i < len(nuova); i++ {
		for j := 0; j < len(nuova[i]); j++ {
			if nuova[i][j] == false {
				fmt.Print(". ")
			} else {
				fmt.Print("# ")
			}
		}
		fmt.Println()
	}
}

func piegax(x int, nuova [][]bool) [][]bool {
	var new [][]bool
	for i := 0; i < len(nuova); i++ { //quante volte
		var in []bool
		for j := 0; j < x-1; j++ {
			if nuova[i][j] == true || nuova[i][len(nuova[i])-1-j] == true {
				in = append(in, true)
			} else {
				in = append(in, false)
			}
		}
		new = append(new, in)
	}
	return new
}

func piegay(y int, nuova [][]bool) [][]bool {
	var new [][]bool
	for k := 0; k < y-1; k++ { //quante volte
		var in []bool
		for j := 0; j < len(nuova[0]); j++ {
			if nuova[k][j] == true || nuova[len(nuova)-1-k][j] == true {
				in = append(in, true)
			} else {
				in = append(in, false)
			}
		}
		new = append(new, in)
	}
	return new
}
*/

package main

func main() {

}
