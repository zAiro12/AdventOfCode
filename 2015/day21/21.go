package main

import (
	"fmt"
	"math"
	"sort"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type entita struct {
	vita     int
	danno    int
	armatura int
}

type oggetti struct {
	costo, val int
}

func main() {
	boss := entita{100, 8, 2}
	giocatore := entita{100, 0, 0}

	armi := []oggetti{{8, 4}, {10, 5}, {25, 6}, {40, 7}, {74, 8}}
	armature := []oggetti{{13, 1}, {31, 2}, {53, 3}, {75, 4}, {102, 5}}
	anelli := []oggetti{{25, 1}, {50, 2}, {100, 3}, {20, 1}, {40, 2}, {80, 3}}

	zairo.StampaA(parteA(giocatore, boss, armi, armature, anelli))
	zairo.StampaB(parteB(giocatore, boss, armi, armature, anelli))
}

func parteA(giocatore, boss entita, armi, armature, anelli []oggetti) int {
	costi := make(map[int]int)

	for i := 0; i < len(armi); i++ {
		costi[armi[i].costo] = armi[i].val
	}

	var chiavi []int
	for k := range costi {
		chiavi = append(chiavi, k)
	}
	sort.Ints(chiavi)
	for i := 0; i < len(chiavi); i++ {
		for j := 0; j < len(armature); j++ {
			if costi[chiavi[i]+armature[j].costo] < costi[chiavi[i]]+armature[j].val {
				costi[chiavi[i]+armature[j].costo] = costi[chiavi[i]] + armature[j].val
			}
		}
	}

	chiavi = []int{}
	for k := range costi {
		chiavi = append(chiavi, k)
	}
	sort.Ints(chiavi)
	for i := 0; i < len(chiavi); i++ {
		for j := 0; j < len(anelli); j++ {
			if costi[chiavi[i]+anelli[j].costo] < costi[chiavi[i]]+anelli[j].val {
				costi[chiavi[i]+anelli[j].costo] = costi[chiavi[i]] + anelli[j].val
			}
		}
	}

	min := math.MaxInt

	for k, v := range costi {
		if k < min && v >= boss.armatura+boss.danno {
			min = k
		}
	}

	return min
}

func parteB(giocatore, boss entita, armi, armature, anelli []oggetti) int {
	costi := make(map[int]int)

	for i := 0; i < len(armi); i++ {
		costi[armi[i].costo] = armi[i].val
	}

	var chiavi []int
	for k := range costi {
		chiavi = append(chiavi, k)
	}
	sort.Ints(chiavi)
	for i := 0; i < len(chiavi); i++ {
		for j := 0; j < len(armature); j++ {
			if costi[chiavi[i]+armature[j].costo] > costi[chiavi[i]]+armature[j].val || costi[chiavi[i]+armature[j].costo] == 0 {
				costi[chiavi[i]+armature[j].costo] = costi[chiavi[i]] + armature[j].val
			}
		}
	}

	chiavi = []int{}
	for k := range costi {
		chiavi = append(chiavi, k)
	}
	sort.Ints(chiavi)
	for i := 0; i < len(chiavi); i++ {
		for j := 0; j < len(anelli); j++ {
			if costi[chiavi[i]+anelli[j].costo] > costi[chiavi[i]]+anelli[j].val || costi[chiavi[i]+anelli[j].costo] == 0 {
				costi[chiavi[i]+anelli[j].costo] = costi[chiavi[i]] + anelli[j].val
			}
		}
	}
	
	chiavi = []int{}
	for k := range costi {
		chiavi = append(chiavi, k)
	}
	sort.Ints(chiavi)
	for i := 0; i < len(chiavi); i++ {
		for j := 0; j < len(anelli); j++ {
			if costi[chiavi[i]+anelli[j].costo] > costi[chiavi[i]]+anelli[j].val || costi[chiavi[i]+anelli[j].costo] == 0 {
				costi[chiavi[i]+anelli[j].costo] = costi[chiavi[i]] + anelli[j].val
			}
		}
	}

	max := 0

	zairo.Log("", costi)

	for k, v := range costi {
		if k > max && v < boss.armatura+boss.danno {
			fmt.Print(k, v, "   ")
			max = k
		}
	}

	return max
}
