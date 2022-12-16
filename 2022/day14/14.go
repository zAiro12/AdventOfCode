package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type dataCascata struct {
	minY  int
	maxY  int
	minX  int
	maxX  int
	mappa map[zairo.Point]bool

	quanteRocce int
}

func main() {
	var cascataA dataCascata = dataCascata{0, 0, 500, 500, make(map[zairo.Point]bool), 0}
	var cascataB dataCascata = dataCascata{0, 0, 500, 500, make(map[zairo.Point]bool), 0}

	sc := bufio.NewScanner(zairo.Input(os.Args[1]))

	for sc.Scan() {
		cascataA.leggiRiga(sc.Text())
		cascataB.leggiRiga(sc.Text())
	}

	zairo.StampaA(parteA(cascataA))
	zairo.StampaB(parteB(&cascataB))
}

func (cascata dataCascata) stampa() {
	for i := cascata.minY; i <= cascata.maxY; i++ {
		for j := cascata.minX; j <= cascata.maxX; j++ {
			if cascata.mappa[zairo.Point{X: j, Y: i}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (cascata *dataCascata) leggiRiga(input string) {
	split := strings.Split(input, " -> ")
	for i := 0; i < len(split)-1; i++ {
		cascata.creaRoccia(split[i], split[i+1])
	}

}

func (cascata *dataCascata) creaRoccia(roccia1, roccia2 string) {
	coordinata1 := strings.Split(roccia1, ",")
	coordinata2 := strings.Split(roccia2, ",")

	x1, _ := strconv.Atoi(coordinata1[0])
	y1, _ := strconv.Atoi(coordinata1[1])
	x2, _ := strconv.Atoi(coordinata2[0])
	y2, _ := strconv.Atoi(coordinata2[1])

	cascata.aggiorna(x1, x2, y1, y2)

	if x1 == x2 {

		for i := y1; i <= y2; i++ {
			cascata.mappa[zairo.Point{X: x1, Y: i}] = true
			cascata.quanteRocce++
		}

		for i := y1; i >= y2; i-- {
			cascata.mappa[zairo.Point{X: x1, Y: i}] = true
			cascata.quanteRocce++
		}

	} else {

		for i := x1; i <= x2; i++ {
			cascata.mappa[zairo.Point{X: i, Y: y1}] = true
			cascata.quanteRocce++
		}

		for i := x1; i >= x2; i-- {
			cascata.mappa[zairo.Point{X: i, Y: y1}] = true
			cascata.quanteRocce++
		}

	}
}

func (cascata *dataCascata) aggiorna(x1, x2, y1, y2 int) {

	if x1 < cascata.minX {
		cascata.maxX = x1
	}
	if x2 < cascata.minX {
		cascata.minX = x2
	}

	if x1 > cascata.maxX {
		cascata.maxX = x1
	}
	if x2 > cascata.maxX {
		cascata.maxX = x2
	}

	if y1 < cascata.minY {
		cascata.minY = y1
	}
	if y2 < cascata.minY {
		cascata.minY = y2
	}

	if y1 > cascata.maxY {
		cascata.maxY = y1
	}
	if y2 > cascata.maxY {
		cascata.maxY = y2
	}
}

func parteA(cascata dataCascata) int {

	granello := zairo.Point{X: 500, Y: 0}
	var c int

a:
	for {
		for {
			if !cascata.mappa[zairo.Point{X: granello.X, Y: granello.Y + 1}] && granello.Y < cascata.maxY {
				granello = zairo.Point{X: granello.X, Y: granello.Y + 1}
				continue
			}
			if !cascata.mappa[zairo.Point{X: granello.X - 1, Y: granello.Y + 1}] && granello.X > cascata.minX && granello.Y < cascata.maxY {
				granello = zairo.Point{X: granello.X - 1, Y: granello.Y + 1}
				continue
			}
			if !cascata.mappa[zairo.Point{X: granello.X + 1, Y: granello.Y + 1}] && granello.X < cascata.maxX && granello.Y < cascata.maxY {
				granello = zairo.Point{X: granello.X + 1, Y: granello.Y + 1}
				continue
			}

			if granello.Y >= cascata.maxY || granello.Y <= cascata.minY || granello.X <= cascata.minX || granello.X >= cascata.maxX {
				break a
			}

			cascata.mappa[granello] = true
			granello = zairo.Point{X: 500, Y: 0}
			break
		}

		c++

	}
	return c
}

func parteB(cascata *dataCascata) int {

	granello := zairo.Point{X: 500, Y: 0}
	var c int

a:
	for {
		for {
			if cascata.mappa[zairo.Point{X: 499, Y: 1}] && cascata.mappa[zairo.Point{X: 500, Y: 1}] && cascata.mappa[zairo.Point{X: 501, Y: 1}] {
				c++
				break a
			}

			if !cascata.mappa[zairo.Point{X: granello.X, Y: granello.Y + 1}] {
				if granello.Y > cascata.maxY {
					cascata.mappa[granello] = true
					granello = zairo.Point{X: 500, Y: 0}
					break
				}

				granello = zairo.Point{X: granello.X, Y: granello.Y + 1}
				continue
			}
			if !cascata.mappa[zairo.Point{X: granello.X - 1, Y: granello.Y + 1}] {
				if granello.X < cascata.minX {
					cascata.minX = granello.X
				}

				if granello.Y > cascata.maxY {
					cascata.mappa[granello] = true
					granello = zairo.Point{X: 500, Y: 0}
					break
				}

				granello = zairo.Point{X: granello.X - 1, Y: granello.Y + 1}
				continue
			}
			if !cascata.mappa[zairo.Point{X: granello.X + 1, Y: granello.Y + 1}] {
				if granello.X > cascata.maxX {
					cascata.maxX = granello.X
				}

				if granello.Y > cascata.maxY {
					cascata.mappa[granello] = true
					granello = zairo.Point{X: 500, Y: 0}
					break
				}

				granello = zairo.Point{X: granello.X + 1, Y: granello.Y + 1}
				continue
			}

			if granello.Y > cascata.maxY {
				cascata.mappa[granello] = true
				granello = zairo.Point{X: 500, Y: 0}
				break
			}

			cascata.mappa[granello] = true
			granello = zairo.Point{X: 500, Y: 0}
			break
		}
		c++

	}
	return c
}
