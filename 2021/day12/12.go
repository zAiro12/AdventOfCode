package main

import (
	"fmt"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	grafo := zairo.NewGrafoString()
	grafo.InputString()
	parteA(grafo)
}

func parteA(grafo zairo.GrafoString) int {
	controllo := make(map[string]bool)
	var percorsi []string
	perocrso(grafo, "start", percorsi, controllo)
	return len(percorsi)
}

func perocrso(grafo zairo.GrafoString, nodo string, percorsi []string, controllo map[string]bool) {
	if nodo == "end" || controllo[nodo] {
		return
	}

	lower := strings.ToLower(nodo)
	if lower == nodo {
		controllo[nodo] = true
	}

	fmt.Print(nodo, " ")
	for i := 0; i < len(grafo[nodo]); i++ {
		perocrso(grafo, grafo[nodo][i], percorsi, controllo)
	}
}
