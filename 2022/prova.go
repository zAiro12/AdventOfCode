package main

import (
	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	g := zairo.NuovoGrafo(8)
	zairo.AggiungiCoppieGrafo(g)
	zairo.StampaGrafo(g)
}
