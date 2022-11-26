package main

import (
	"fmt"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	g := zairo.NuovoGrafo(8)
	zairo.AggiungiCoppieGrafo(g)
	zairo.StampaGrafo(g)

func Trova(grafo *zairo.Grafo, x, y int) bool {
	if x > grafo.Len || x < 0 {
		return false
	}
	appoggio := grafo.Adiacenti[x]
	for appoggio != nil {
		if appoggio.Val == y {
			return true
		}
		appoggio = appoggio.Next
	}
	return false
}
