package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {

	//parte A
	var istruzione string
	var split []string
	var input, x, potenzaSegnale, ciclo int
	x = 1

	//parte B
	var CRT [6][40]bool

	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)
	for sc.Scan() {

		istruzione = sc.Text()

		ciclo++
		potenzaSegnale += parteA(ciclo, x)
		CRT = parteB(ciclo, x, CRT)

		if istruzione != "noop" {
			split = strings.Split(istruzione, " ")
			input, _ = strconv.Atoi(split[1])

			ciclo++
			potenzaSegnale += parteA(ciclo, x)
			CRT = parteB(ciclo, x, CRT)
			x += input
		}

	}
	zairo.StampaA(potenzaSegnale)
	zairo.StampaB()
	stampaCRT(CRT)
}

func parteA(ciclo, x int) int {
	if ciclo == 20 || ciclo == 60 || ciclo == 100 || ciclo == 140 || ciclo == 180 || ciclo == 220 {
		return ciclo * x
	}
	return 0
}

func parteB(ciclo, x int, CRT [6][40]bool) [6][40]bool {
	cursore := (ciclo - 1) % 40
	riga := (ciclo - 1) / 40
	if x == cursore-1 || x == cursore || x == cursore+1 {
		CRT[riga][cursore] = true
	}

	return CRT
}

func stampaCRT(CRT [6][40]bool) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			if CRT[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
