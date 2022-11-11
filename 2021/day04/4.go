package main

import "fmt"

type matrix struct {
	x [5][5]int
}

func main() {
	var gen []int
	var in int
	var matrice []matrix
	for i := 0; i < 100; i++ {
		fmt.Scan(&in)
		gen = append(gen, in)
	}
	for i := 0; ; i++ {
		mat, fine := riempi(matrice[i])
		matrice = append(matrice, mat)
		if fine {
			break
		}
	}

	fmt.Println(matrice)
}

func riempi(matrice matrix) (matrix, bool) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			var in int
			_, err := fmt.Scan(&in)
			if err != nil {
				return matrice, true
			}
			matrice.x[i][j] = in
		}
	}
	return matrice, false
}
