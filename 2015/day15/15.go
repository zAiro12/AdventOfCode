package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type biscotto struct {
	nome       string
	capacita   int
	durabilita int
	sapore     int
	texture    int
	calorie    int
}

func main() {
	biscotti := make(map[int]biscotto)

	file, _ := os.Open(os.Args[1])
	defer file.Close()

	sc := bufio.NewScanner(file)

	var appoggio biscotto
	var i int
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &appoggio.nome, &appoggio.capacita, &appoggio.durabilita, &appoggio.sapore, &appoggio.texture, &appoggio.calorie)

		appoggio.nome = strings.Trim(appoggio.nome, ":")
		biscotti[i] = appoggio
		i++
	}

	zairo.StampaA(parteA(biscotti))
	zairo.StampaB(parteB(biscotti))
}

func parteA(biscotti map[int]biscotto) int {

	var max, ris int
	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			for k := 0; k < 101; k++ {
				for z := 0; z < 101; z++ {
					if i+j+k+z == 100 {
						ris = calcolaBiscotti(i, j, k, z, biscotti)
						if ris > max {
							max = ris
						}
					}
				}
			}
		}
	}

	return max
}

func calcolaBiscotti(i, j, k, z int, biscotti map[int]biscotto) int {
	var risultato []int
	var tempRis int

	tempRis = i*biscotti[0].capacita + j*biscotti[1].capacita + k*biscotti[2].capacita + z*biscotti[3].capacita
	if tempRis < 0 {
		return 0
	} else {
		risultato = append(risultato, tempRis)
	}

	tempRis = i*biscotti[0].durabilita + j*biscotti[1].durabilita + k*biscotti[2].durabilita + z*biscotti[3].durabilita
	if tempRis < 0 {
		return 0
	} else {
		risultato = append(risultato, tempRis)
	}

	tempRis = i*biscotti[0].sapore + j*biscotti[1].sapore + k*biscotti[2].sapore + z*biscotti[3].sapore
	if tempRis < 0 {
		return 0
	} else {
		risultato = append(risultato, tempRis)
	}

	tempRis = i*biscotti[0].texture + j*biscotti[1].texture + k*biscotti[2].texture + z*biscotti[3].texture
	if tempRis < 0 {
		return 0
	} else {
		risultato = append(risultato, tempRis)
	}

	return risultato[0] * risultato[1] * risultato[2] * risultato[3]
}

func parteB(biscotti map[int]biscotto) int {
	var max, ris int
	for i := 0; i < 101; i++ {
		for j := 0; j < 101; j++ {
			for k := 0; k < 101; k++ {
				for z := 0; z < 101; z++ {
					if i+j+k+z == 100 {
						ris = calcolaCalorie(i, j, k, z, biscotti)
						if ris > max {
							max = ris
						}
					}
				}
			}
		}
	}

	return max
}

func calcolaCalorie(i, j, k, z int, biscotti map[int]biscotto) int {

	ris := i*biscotti[0].calorie + j*biscotti[1].calorie + k*biscotti[2].calorie + z*biscotti[3].calorie

	if ris != 500 {
		return 0
	}
	return calcolaBiscotti(i, j, k, z, biscotti)
}
