package main

import (
	"fmt"
	"math"
)

type posizione struct {
	x, y int
}

func main() {
	var p1, p2 []posizione
	var c int

	for {
		var in int
		var punto posizione
		var inutile string
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		punto.x = in
		fmt.Scan(&in)
		punto.y = in
		p1 = append(p1, punto)

		fmt.Scan(&inutile)

		fmt.Scan(&in)
		punto.x = in
		fmt.Scan(&in)
		punto.y = in
		p2 = append(p2, punto)
	}

	p1, p2 = line(p1, p2)
	griglia := punti(p1, p2)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if griglia[i][j] >= 2 {
				c++
			}
		}
	}
	fmt.Println(c)
}

func line(p1, p2 []posizione) ([]posizione, []posizione) {
	var giusto1, giusto2 []posizione
	for i := 0; i < len(p1); i++ {
		if p1[i].x == p2[i].x || p1[i].y == p2[i].y {
			giusto1 = append(giusto1, p1[i])
			giusto2 = append(giusto2, p2[i])
		}
	}
	return giusto1, giusto2

}

func punti(p1, p2 []posizione) [1000][1000]int {
	var griglia [1000][1000]int
	for i := 0; i < len(p1); i++ {
		if p1[i].x == p2[i].x {
			min := int(math.Min(float64(p1[i].y), float64(p2[i].y)))
			max := int(math.Max(float64(p1[i].y), float64(p2[i].y)))
			for j := min; j <= max; j++ {
				griglia[p1[i].x][j]++
			}
		}
		if p1[i].y == p2[i].y {
			min := int(math.Min(float64(p1[i].x), float64(p2[i].x)))
			max := int(math.Max(float64(p1[i].x), float64(p2[i].x)))
			for j := min; j <= max; j++ {
				griglia[j][p1[i].y]++
			}
		}
	}
	return griglia
}
