package main

import (
	"fmt"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	var word string
	fmt.Scan(&word)

	zairo.StampaA(parteA(word))
	zairo.StampaB(parteB(word))
}

func parteA(word string) int {
	var c int

	for i := 0; i < len(word)-4; i++ {
		if isDiverse(word[i], word[i+1], word[i+2], word[i+3]) {
			return c + 4
		}

		c++
	}

	return 0
}

func isDiverse(uno, due, tre, quattro byte) bool {
	ricorrenze := make(map[byte]bool)

	ricorrenze[uno] = true
	ricorrenze[due] = true
	ricorrenze[tre] = true
	ricorrenze[quattro] = true

	return len(ricorrenze) == 4
}

func parteB(word string) int {
	var c int

	for i := 0; len(word) != 0; i++ {
		if isDiverseB(word[:14]) {
			return c + 14

		} else {
			word = word[1:]
		}

		c++
	}
	return 0
}

func isDiverseB(lettere string) bool {
	ricorrenze := make(map[byte]bool)

	for i := 0; i < len(lettere); i++ {
		ricorrenze[lettere[i]] = true
	}

	return len(ricorrenze) == 14
}
