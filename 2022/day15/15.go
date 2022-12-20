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

	insieme := dato{make(map[zairo.Point]bool), make(map[zairo.Point]zairo.Point), 0, 0, 0, 0}

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
	zairo.StampaB()
}

func (insieme *dato) aggiornaMatrice(sensore, beacon zairo.Point) {

	insieme.matrice[sensore] = true
	insieme.matrice[beacon] = true
	insieme.legami[sensore] = beacon

	minX := zairo.Min(sensore.X, beacon.X)
	maxX := zairo.Max(sensore.X, beacon.X)
	minY := zairo.Min(sensore.Y, beacon.Y)
	maxY := zairo.Max(sensore.Y, beacon.Y)

	if minX < insieme.minX {
		insieme.minX = minX
	}
	if maxX > insieme.maxX {
		insieme.maxX = maxX
	}

	if minY < insieme.minY {
		insieme.minY = minY
	}
	if maxY > insieme.maxY {
		insieme.maxY = maxY
	}
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

func (insieme dato) stampaRiga(y int) {
	var c int
	for i := insieme.minX; i <= insieme.maxX; i++ {

		if insieme.matrice[zairo.Point{X: i, Y: y}] {
			fmt.Print("#")
			c++
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println(c)
}

func (insieme *dato) parteA() int {
	var distanza int
	for sensor, beacon := range insieme.legami {
		zairo.Log("ciclo", sensor, beacon)
		distanza = calcolaDistanza(sensor, beacon)
		insieme.creaMacchia(sensor, distanza)
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
	zairo.Log("distanza", distanza)
	var tmp1, tmp2, tmp3, tmp4 zairo.Point
	for i := 0; i <= distanza; i++ {
		for j := 0; j <= distanza; j++ {
			tmp1 = zairo.Point{X: start.X + j, Y: start.Y + i}
			tmp2 = zairo.Point{X: start.X - j, Y: start.Y + i}
			tmp3 = zairo.Point{X: start.X - j, Y: start.Y - i}
			tmp4 = zairo.Point{X: start.X + j, Y: start.Y - i}

			if calcolaDistanza(tmp1, start) <= distanza {
				insieme.matrice[tmp1] = true
				if tmp1.X > insieme.maxX {
					insieme.maxX = tmp1.X
				} else if tmp1.Y > insieme.maxY {
					insieme.maxY = tmp1.Y
				}
			}
			if calcolaDistanza(tmp2, start) <= distanza {
				insieme.matrice[tmp2] = true
				if tmp2.X > insieme.maxX {
					insieme.maxX = tmp2.X
				} else if tmp2.Y < insieme.minY {
					insieme.minY = tmp2.Y
				}
			}
			if calcolaDistanza(tmp3, start) <= distanza {
				insieme.matrice[tmp3] = true
				if tmp3.X < insieme.minX {
					insieme.minX = tmp3.X
				} else if tmp3.Y < insieme.minX {
					insieme.minY = tmp3.Y
				}
			}
			if calcolaDistanza(tmp4, start) <= distanza {
				insieme.matrice[tmp4] = true
				if tmp4.X < insieme.minX {
					insieme.minX = tmp4.X
				} else if tmp4.Y > insieme.maxY {
					insieme.maxY = tmp4.Y
				}
			}
		}
	}
}
