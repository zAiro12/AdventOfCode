package main

import (
	"fmt"

	"github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	var opponent, mio string
	var punteggioA, punteggioB int
	for {
		_, err := fmt.Scan(&opponent)

		if err != nil {
			break
		}
		fmt.Scan(&mio)
		punteggioA += parteA(opponent, mio)
		punteggioB += parteB(opponent, mio)
	}

	utile.StampaA(punteggioA)
	utile.StampaB(punteggioB)
}

func parteA(opponent, mio string) int {
	switch opponent {
	case "A":
		if mio == "X" {
			return 4
		} else if mio == "Y" {
			return 8
		} else {
			return 3
		}

	case "B":
		if mio == "X" {
			return 1
		} else if mio == "Y" {
			return 5
		} else {
			return 9
		}

	case "C":
		if mio == "X" {
			return 7
		} else if mio == "Y" {
			return 2
		} else {
			return 6
		}

	}
	return 0
}

func parteB(opponent, mio string) int {
	switch opponent {
	case "A":
		if mio == "X" {
			return 3
		} else if mio == "Y" {
			return 1 + 3
		} else {
			return 2 + 6
		}

	case "B":
		if mio == "X" {
			return 1
		} else if mio == "Y" {
			return 2 + 3
		} else {
			return 3 + 6
		}

	case "C":
		if mio == "X" {
			return 2
		} else if mio == "Y" {
			return 3 + 3
		} else {
			return 1 + 6
		}

	}
	return 0
}
