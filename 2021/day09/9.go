package main

import "fmt"

func main() {
	var c int
	var matrice [100][]int
	for i := 0; ; i++ {
		var in string
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		newin := []rune(in)
		for j := 0; j < len(newin); j++ {
			matrice[i] = append(matrice[i], int(newin[j])-48)
		}
	}
	for i := 0; i < 100; i++ {
		for j := 0; j < len(matrice[i]); j++ {
			//controllo lato sx
			if i == 0 {
				if j == 0 {
					if matrice[i][j] < matrice[i][j+1] && matrice[i][j] < matrice[i+1][j] {
						c += matrice[i][j] + 1
					}
				} else if j == len(matrice[i])-1 {
					if matrice[i][j] < matrice[i+1][j] && matrice[i][j] < matrice[i][j-1] {
						c += matrice[i][j] + 1
					}
				} else {
					if matrice[i][j] < matrice[i+1][j] && matrice[i][j] < matrice[i][j-1] && matrice[i][j] < matrice[i][j+1] {
						c += matrice[i][j] + 1
					}
				}
				//controllo lato dx
			} else if i == len(matrice)-1 {
				if j == 0 {
					if matrice[i][j] < matrice[i][j+1] && matrice[i][j] < matrice[i-1][j] {
						c += matrice[i][j] + 1
					}
				} else if j == len(matrice[i])-1 {
					if matrice[i][j] < matrice[i-1][j] && matrice[i][j] < matrice[i][j-1] {
						c += matrice[i][j] + 1
					}
				} else {
					if matrice[i][j] < matrice[i-1][j] && matrice[i][j] < matrice[i][j-1] && matrice[i][j] < matrice[i][j+1] {
						c += matrice[i][j] + 1
					}
				}
				//controllo lato top
			} else if j == 0 {
				if matrice[i][j] < matrice[i-1][j] && matrice[i][j] < matrice[i+1][j] && matrice[i][j] < matrice[i][j+1] {
					c += matrice[i][j] + 1
				}
				//controllo lato bot
			} else if j == len(matrice[i])-1 {
				if matrice[i][j] < matrice[i-1][j] && matrice[i][j] < matrice[i+1][j] && matrice[i][j] < matrice[i][j-1] {
					c += matrice[i][j] + 1
				}
				//caso generale
			} else {
				if matrice[i][j] < matrice[i+1][j] && matrice[i][j] < matrice[i-1][j] && matrice[i][j] < matrice[i][j+1] && matrice[i][j] < matrice[i][j-1] {
					c += matrice[i][j] + 1
				}
			}
		}
	}
	fmt.Println(c)
}
