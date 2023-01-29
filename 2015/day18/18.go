package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type point struct {
	y, x int
}

type griglia struct {
	griglia    map[point]bool
	maxX, maxY int
}

func main() {
	file, _ := os.Open(os.Args[1])
	sc := bufio.NewScanner(file)
	defer file.Close()

	grigliaA := griglia{make(map[point]bool), 0, 0}
	grigliaB := griglia{make(map[point]bool), 0, 0}
	var split []string

	for sc.Scan() {
		split = strings.Split(sc.Text(), "")
		grigliaB.maxX = len(split)

		for x := 0; x < len(split); x++ {
			if split[x] == "#" {
				grigliaA.griglia[point{grigliaA.maxY, x}] = true
				grigliaB.griglia[point{grigliaB.maxY, x}] = true

			}
		}
		grigliaA.maxY++
		grigliaB.maxY++
	}
	grigliaA.maxX = len(split)
	grigliaB.maxX = len(split)

	grigliaB.griglia[point{0, 0}] = true
	grigliaB.griglia[point{grigliaB.maxX - 1, 0}] = true
	grigliaB.griglia[point{0, grigliaB.maxY - 1}] = true
	grigliaB.griglia[point{grigliaB.maxX - 1, grigliaB.maxY - 1}] = true

	zairo.StampaA(ParteA(grigliaA, 100))
	zairo.StampaB(ParteB(grigliaB, 100))
}

func ParteA(griglia griglia, step int) int {
	for i := 0; i < step; i++ {
		matriceClone := make(map[point]bool)

		for y := 0; y < griglia.maxY; y++ {
			for x := 0; x < griglia.maxX; x++ {

				matriceClone[point{y, x}] = false
			}
		}

		for y := 0; y < griglia.maxY; y++ {
			for x := 0; x < griglia.maxX; x++ {

				AggiornaPunto(point{y, x}, griglia.griglia, matriceClone)
			}
		}
		griglia.griglia = matriceClone
	}

	var counter int
	for y := 0; y < griglia.maxY; y++ {
		for x := 0; x < griglia.maxX; x++ {
			if griglia.griglia[point{y, x}] {
				counter++
			}
		}
	}
	return counter
}

func AggiornaPunto(daAnalizzare point, griglia, clone map[point]bool) {
	var counter int

	if griglia[point{daAnalizzare.y + 1, daAnalizzare.x}] {
		counter++
	}
	if griglia[point{daAnalizzare.y + 1, daAnalizzare.x + 1}] {
		counter++
	}
	if griglia[point{daAnalizzare.y, daAnalizzare.x + 1}] {
		counter++
	}
	if griglia[point{daAnalizzare.y - 1, daAnalizzare.x + 1}] {
		counter++
	}
	if griglia[point{daAnalizzare.y - 1, daAnalizzare.x}] {
		counter++
	}
	if griglia[point{daAnalizzare.y - 1, daAnalizzare.x - 1}] {
		counter++
	}
	if griglia[point{daAnalizzare.y, daAnalizzare.x - 1}] {
		counter++
	}
	if griglia[point{daAnalizzare.y + 1, daAnalizzare.x - 1}] {
		counter++
	}

	if griglia[daAnalizzare] && (counter == 2 || counter == 3) {
		clone[daAnalizzare] = true
	}
	if !griglia[daAnalizzare] && counter == 3 {
		clone[daAnalizzare] = true
	}
}

func ParteB(griglia griglia, step int) int {
	for i := 0; i < step; i++ {
		matriceClone := make(map[point]bool)

		for y := 0; y < griglia.maxY; y++ {
			for x := 0; x < griglia.maxX; x++ {

				if (x == 0 && y == 0) || (x == griglia.maxX-1 && y == 0) || (x == 0 && y == griglia.maxY-1) || (x == griglia.maxX-1 && y == griglia.maxY-1) {
					matriceClone[point{y, x}] = true
				} else {
					matriceClone[point{y, x}] = false
				}

			}
		}

		for y := 0; y < griglia.maxY; y++ {
			for x := 0; x < griglia.maxX; x++ {

				AggiornaPunto(point{y, x}, griglia.griglia, matriceClone)
			}
		}

		griglia.griglia = matriceClone

	}

	var counter int
	for y := 0; y < griglia.maxY; y++ {
		for x := 0; x < griglia.maxX; x++ {
			if griglia.griglia[point{y, x}] {
				counter++
			}
		}
	}
	return counter
}

func StampaGriglia(griglia griglia) {
	for y := 0; y < griglia.maxY; y++ {
		for x := 0; x < griglia.maxX; x++ {

			if griglia.griglia[point{y, x}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
