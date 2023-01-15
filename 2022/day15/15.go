package main

import (
	"bufio"
	"fmt"
	"os"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type dato struct {
	matrice map[zairo.Point]bool
	legami  map[zairo.Point]zairo.Point

	maxX int
	maxY int
	minX int
	minY int
}

func main() {

	insieme := dato{make(map[zairo.Point]bool), make(map[zairo.Point]zairo.Point), 0, 20, 0, 20}

	sc := bufio.NewScanner(zairo.Input(os.Args[1]))
	var sensore, beacon zairo.Point
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensore.X, &sensore.Y, &beacon.X, &beacon.Y)

		insieme.aggiornaMatrice(sensore, beacon)

	}

	//insieme.stampaMatrice()

	zairo.Log("min/max", insieme.minX, insieme.maxX, insieme.minY, insieme.maxY)
	zairo.StampaA(insieme.parteA())
	//insieme.stampaMatrice()
	zairo.StampaB(insieme.parteB())
}

func (insieme *dato) aggiornaMatrice(sensore, beacon zairo.Point) {

	insieme.matrice[sensore] = true
	insieme.matrice[beacon] = true
	insieme.legami[sensore] = beacon

}

func (insieme dato) stampaMatrice() {
	for y := insieme.minY; y <= insieme.maxY; y++ {
		for x := insieme.minX; x <= insieme.maxX; x++ {
			if insieme.matrice[zairo.Point{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (insieme *dato) parteA() int {
	var distanza int
	for sensor, beacon := range insieme.legami {
		distanza = calcolaDistanza(sensor, beacon)

		insieme.creaRiga(sensor, distanza)
	}
	var counter int
	for i := insieme.minX; i <= insieme.maxX; i++ {
		if insieme.matrice[zairo.Point{X: i, Y: 2000000}] {
			counter++
		}
	}
	return counter - 1
}

func calcolaDistanza(puntoA, puntoB zairo.Point) int {
	return zairo.ABS(puntoA.X-puntoB.X) + zairo.ABS(puntoA.Y-puntoB.Y)
}

func (insieme *dato) creaMacchia(start zairo.Point, distanza int) {
	var tmp1, tmp2, tmp3, tmp4 zairo.Point
	for i := 0; i <= distanza; i++ {
		for j := 0; j <= distanza; j++ {
			tmp1 = zairo.Point{X: start.X + j, Y: start.Y + i}
			tmp2 = zairo.Point{X: start.X - j, Y: start.Y + i}
			tmp3 = zairo.Point{X: start.X - j, Y: start.Y - i}
			tmp4 = zairo.Point{X: start.X + j, Y: start.Y - i}

			if calcolaDistanza(tmp1, start) <= distanza {
				insieme.matrice[tmp1] = true

			}
			if calcolaDistanza(tmp2, start) <= distanza {
				insieme.matrice[tmp2] = true

			}
			if calcolaDistanza(tmp3, start) <= distanza {
				insieme.matrice[tmp3] = true

			}
			if calcolaDistanza(tmp4, start) <= distanza {
				insieme.matrice[tmp4] = true

			}
		}
	}
}

func (insieme *dato) creaRiga(start zairo.Point, distanza int) {
	r := zairo.ABS(2000000 - start.Y)

	if distanza >= r {

		for i := start.X - (distanza - r); i <= start.X+(distanza-r); i++ {
			insieme.matrice[zairo.Point{X: i, Y: 2000000}] = true
		}

		if start.X-(distanza-r) < insieme.minX {
			insieme.minX = start.X - (distanza - r)
		}

		if start.X+(distanza-r) > insieme.maxX {
			insieme.maxX = start.X + (distanza - r)
		}
	}

}

func (insieme *dato) parteB() int {
	insieme.minX, insieme.minY = 0, 0
	//insieme.maxX, insieme.maxY = 20, 20
	insieme.maxX, insieme.maxY = 4000000, 4000000

	var distanza int
	for sensor, beacon := range insieme.legami {
		distanza = calcolaDistanza(sensor, beacon)

		// insieme.creaMacchia(sensor, distanza)
		_ = distanza
	}

	for y := insieme.minY; y <= insieme.maxY; y++ {
		for x := insieme.minX; x <= insieme.maxX; x++ {
			if insieme.matrice[zairo.Point{X: x, Y: y}] {
				return x*4000000 + y
			}
		}
	}

	return 0
}
