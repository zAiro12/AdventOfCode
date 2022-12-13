package main

import (
	"bufio"
	"os"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	matrice := make(map[zairo.Point]string)

	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)

	var in string
	var split []string

	var y, x int

	for sc.Scan() {
		in = sc.Text()
		split = strings.Split(in, "")

		for x = 0; x < len(split); x++ {
			matrice[zairo.Point{X: x, Y: y}] = split[x]
		}
		y++
	}

	zairo.StampaMappa(matrice, x, y)
	zairo.StampaA()
	zairo.StampaB()
}
