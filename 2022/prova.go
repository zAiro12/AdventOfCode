package main

import (
	"fmt"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	g := zairo.NuovoGrafo(8)
	zairo.AggiungiCoppieGrafo(g)
	zairo.StampaGrafo(g)
	fmt.Println(Trova(g, 0, 1))
}

func Trova(grafo *zairo.Grafo, x, y int) bool {
	appoggio := grafo.Adiacenti[x]
	for appoggio != nil {
		if appoggio.Val == y {
			return true
		}
		appoggio = appoggio.Next
	}
	return false
}
