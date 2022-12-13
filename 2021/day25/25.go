package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {

	//nella mappa 0: non c'è, 1: destra, 2: giù
	mappa := make(map[zairo.Point]int)
	var appoggio zairo.Point
	var in string
	var split []string

	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		in = sc.Text()
		split = strings.Split(in, "")

		for i := 0; i < len(split); i++ {
			appoggio.X = i

			if split[i] == "." {
				mappa[appoggio] = 0

			} else if split[i] == ">" {
				mappa[appoggio] = 1

			} else {
				mappa[appoggio] = 2
			}
		}
		appoggio.Y++
		appoggio.X++
	}
	//appoggio diventa la lunghezza della riga della mappa

	stampaMappa(mappa, appoggio)

	zairo.StampaA(parteA(mappa, appoggio))
	zairo.StampaB()
}

func stampaMappa(mappa map[zairo.Point]int, lunghezza zairo.Point) {
	for j := 0; j < lunghezza.Y; j++ {
		for i := 0; i < lunghezza.X; i++ {
			if mappa[zairo.Point{X: i, Y: j}] == 0 {
				fmt.Print(".")

			} else if mappa[zairo.Point{X: i, Y: j}] == 1 {
				fmt.Print(">")

			} else {
				fmt.Print("v")
			}
		}
		fmt.Println()
	}
}

func parteA(mappa map[zairo.Point]int, lunghezza zairo.Point) (step int) {

	var flag1, flag2 bool

	// for {
	flag1, mappa = muoviDestra(mappa, lunghezza)
	flag2, mappa = muoviGiu(mappa, lunghezza)

	if flag1 && flag2 {
		step++
	} else {
		// break
	}
	zairo.Log("step: ", step)
	// }

	stampaMappa(mappa, lunghezza)
	return step
}

func muoviDestra(mappa map[zairo.Point]int, lunghezza zairo.Point) (bool, map[zairo.Point]int) {
	var mosso bool

	for k, v := range mappa {
		if v == 1 {

			if k.X < lunghezza.X && (mappa[zairo.Point{X: k.X + 1, Y: k.Y}] == 0) {
				mappa[k] = 0
				mappa[zairo.Point{X: k.X + 1, Y: k.Y}] = 1

			} else if k.X > lunghezza.X && mappa[zairo.Point{X: 0, Y: k.Y}] == 0 {
				mappa[k] = 0
				mappa[zairo.Point{X: 0, Y: k.Y}] = 1
			}

			mosso = true
		}
	}
	return mosso, mappa
}

func muoviGiu(mappa map[zairo.Point]int, lunghezza zairo.Point) (bool, map[zairo.Point]int) {
	var mosso bool

	for k, v := range mappa {

		if v == 2 {

			if k.Y < lunghezza.Y && (mappa[zairo.Point{X: k.X, Y: k.Y + 1}] == 0) {
				mappa[k] = 0
				mappa[zairo.Point{X: k.X + 1, Y: k.Y}] = 2

			} else if k.Y > lunghezza.Y && mappa[zairo.Point{X: k.X, Y: 0}] == 0 {
				mappa[k] = 0
				mappa[zairo.Point{X: k.X, Y: 0}] = 2
			}

			mosso = true
		}

	}
	return mosso, mappa
}
