package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type insieme struct {
	testa      zairo.Point
	coda       zairo.Point
	snake      [10]zairo.Point
	posizioniA map[zairo.Point]bool
	posizioniB map[zairo.Point]bool
}

func main() {
	posizioni := insieme{zairo.Point{X: 0, Y: 0}, zairo.Point{X: 0, Y: 0}, [10]zairo.Point{}, make(map[zairo.Point]bool), make(map[zairo.Point]bool)}

	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)

	var in string
	var split []string
	var num int

	for sc.Scan() {
		in = sc.Text()
		split = strings.Split(in, " ")
		num, _ = strconv.Atoi(split[1])

		posizioni = parteA(posizioni, split[0], num)
		posizioni = parteB(posizioni, split[0], num)
	}

	zairo.StampaA(len(posizioni.posizioniA))
	zairo.StampaB(len(posizioni.posizioniB))
}

func parteA(posizioni insieme, direzione string, volte int) insieme {

	for i := 0; i < volte; i++ {
		switch direzione {
		case "U":
			posizioni.testa.Y++

		case "D":
			posizioni.testa.Y--

		case "R":
			posizioni.testa.X++

		case "L":
			posizioni.testa.X--
		}

		posizioni.coda = muoviCoda(posizioni.testa, posizioni.coda)
		posizioni.posizioniA[posizioni.coda] = true
	}
	return posizioni
}

func muoviCoda(testa, coda zairo.Point) zairo.Point {

	// Nord Ovest
	if (testa.Y-coda.Y == 2 && coda.X-testa.X == 1) || (testa.Y-coda.Y == 2 && coda.X-testa.X == 2) || (testa.Y-coda.Y == 1 && coda.X-testa.X == 2) {
		coda.X--
		coda.Y++
	}

	// Nord
	if testa.X-coda.X == 0 && testa.Y-coda.Y == 2 {
		coda.Y++
	}

	// Nord Est
	if (testa.Y-coda.Y == 2 && testa.X-coda.X == 1) || (testa.Y-coda.Y == 2 && testa.X-coda.X == 2) || (testa.Y-coda.Y == 1 && testa.X-coda.X == 2) {
		coda.X++
		coda.Y++
	}

	// Est
	if testa.Y-coda.Y == 0 && testa.X-coda.X == 2 {
		coda.X++
	}

	// Sud Est
	if (coda.Y-testa.Y == 2 && testa.X-coda.X == 1) || (coda.Y-testa.Y == 2 && testa.X-coda.X == 2) || (coda.Y-testa.Y == 1 && testa.X-coda.X == 2) {
		coda.X++
		coda.Y--
	}

	// Sud
	if testa.X-coda.X == 0 && coda.Y-testa.Y == 2 {
		coda.Y--
	}

	// Sud Ovest
	if (coda.Y-testa.Y == 2 && coda.X-testa.X == 1) || (coda.Y-testa.Y == 2 && coda.X-testa.X == 2) || (coda.Y-testa.Y == 1 && coda.X-testa.X == 2) {
		coda.X--
		coda.Y--
	}

	// Ovest
	if testa.Y-coda.Y == 0 && coda.X-testa.X == 2 {
		coda.X--
	}

	return coda
}

func parteB(posizioni insieme, direzione string, volte int) insieme {

	for i := 0; i < volte; i++ {
		switch direzione {
		case "U":
			posizioni.snake[0].Y++

		case "D":
			posizioni.snake[0].Y--

		case "R":
			posizioni.snake[0].X++

		case "L":
			posizioni.snake[0].X--
		}
		for j := 0; j < 9; j++ {
			posizioni.snake[j+1] = muoviCoda(posizioni.snake[j], posizioni.snake[j+1])
		}
		posizioni.posizioniB[posizioni.snake[9]] = true
	}

	return posizioni
}
