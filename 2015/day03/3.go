package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	//fmt.Println(parteA())
	fmt.Println(parteB())
}

func parteA() int {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanRunes)
	mappa := make(map[point]bool)
	punto := point{0, 0}
	mappa[punto] = true
	for sc.Scan() {
		switch sc.Text() {
		case "^":
			punto.y++
		case ">":
			punto.x++
		case "v":
			punto.y--
		case "<":
			punto.x--
		}
		mappa[punto] = true
	}
	return len(mappa)
}

func parteB() int {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanRunes)
	mappa := make(map[point]bool)

	santa := point{0, 0}
	robot := point{0, 0}
	mappa[santa] = true

	flag := false

	for sc.Scan() {
		if flag {
			switch sc.Text() {
			case "^":
				santa.y++
			case ">":
				santa.x++
			case "v":
				santa.y--
			case "<":
				santa.x--
			}
			mappa[santa] = true
		} else {
			switch sc.Text() {
			case "^":
				robot.y++
			case ">":
				robot.x++
			case "v":
				robot.y--
			case "<":
				robot.x--
			}
			mappa[robot] = true
		}
		flag = !flag
	}

	return len(mappa)
}
