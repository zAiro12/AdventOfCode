package main

import (
	"bufio"
	"fmt"
	"os"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	file, _ := os.Open(os.Args[1])
	sc := bufio.NewScanner(file)
	defer file.Close()

	daCercare := make(map[string]int)

	daCercare["children:"] = 3
	daCercare["cats:"] = 7
	daCercare["samoyeds:"] = 2
	daCercare["pomeranians:"] = 3
	daCercare["akitas:"] = 0
	daCercare["vizslas:"] = 0
	daCercare["goldfish:"] = 5
	daCercare["trees:"] = 3
	daCercare["cars:"] = 2
	daCercare["perfumes:"] = 1

	var risA, risB int
	for sc.Scan() {
		tempA := parteA(sc.Text(), daCercare)
		tempB := parteB(sc.Text(), daCercare)
		if tempA > 0 {
			risA = tempA
		}
		if tempB > 0 && tempB != risA {
			risB = tempB
		}
	}

	zairo.StampaA(risA)
	zairo.StampaB(risB)
}

func parteA(testo string, daCercare map[string]int) int {
	var chiave1, chiave2, chiave3 string
	var valore1, valore2, valore3 int
	var index int
	fmt.Sscanf(testo, "Sue %d: %s %d, %s %d, %s %d", &index, &chiave1, &valore1, &chiave2, &valore2, &chiave3, &valore3)

	if daCercare[chiave1] != valore1 || daCercare[chiave2] != valore2 || daCercare[chiave3] != valore3 {
		return 0
	}
	return index
}

func parteB(testo string, daCercare map[string]int) int {
	var chiavi [3]string
	var valori [3]int
	var index int
	fmt.Sscanf(testo, "Sue %d: %s %d, %s %d, %s %d", &index, &chiavi[0], &valori[0], &chiavi[1], &valori[1], &chiavi[2], &valori[2])

	for i := 0; i < 3; i++ {

		switch chiavi[i] {

		case "cats:", "trees:":
			if daCercare[chiavi[i]] > valori[i] {
				return 0
			}

		case "pomeranians:", "goldfish:":
			if daCercare[chiavi[i]] < valori[i] {
				return 0
			}

		default:
			if daCercare[chiavi[i]] != valori[i] {
				return 0
			}
		}
	}

	return index
}
