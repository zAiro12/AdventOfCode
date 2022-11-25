package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var matrice [][]int
	var num int
	var in string
	var split []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		in = sc.Text()
		split = strings.Split(in, "")
		var appoggio []int
		for i := 0; i < len(split); i++ {
			num, _ = strconv.Atoi(split[i])
			appoggio = append(appoggio, num)
		}
		matrice = append(matrice, appoggio)
	}
	stampa(matrice)

	
}

func stampa(matrice [][]int) {
	for i := 0; i < len(matrice); i++ {
		for j := 0; j < len(matrice[i]); j++ {
			fmt.Print(matrice[i][j])
		}
		fmt.Println()
	}
}
