package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	luci := make(map[point]bool)
	var in string
	var split, str []string
	var x, y int
	var da, a point
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		in = sc.Text()
		split = strings.Split(in, " ")

		if len(split) == 4 {
			str = strings.Split(split[1], ",")
			x, _ = strconv.Atoi(str[0])
			y, _ = strconv.Atoi(str[1])
			da = point{x, y}
			str = strings.Split(split[3], ",")
			x, _ = strconv.Atoi(str[0])
			y, _ = strconv.Atoi(str[1])
			a = point{x, y}
			luci = toggle(luci, da, a)
		} else {
			str = strings.Split(split[2], ",")
			x, _ = strconv.Atoi(str[0])
			y, _ = strconv.Atoi(str[1])
			da = point{x, y}
			str = strings.Split(split[4], ",")
			x, _ = strconv.Atoi(str[0])
			y, _ = strconv.Atoi(str[1])
			a = point{x, y}
			if split[1] == "on" {
				luci = turnOn(luci, da, a)
			} else {
				luci = turnOff(luci, da, a)
			}
		}
	}
	fmt.Println(stampatrue(luci))
}

func turnOn(luci map[point]bool, da, a point) map[point]bool {
	for i := da.x; i <= a.x; i++ {
		for j := da.y; j <= a.y; j++ {
			luci[point{i, j}] = true
		}
	}
	return luci
}

func turnOff(luci map[point]bool, da, a point) map[point]bool {
	for i := da.x; i <= a.x; i++ {
		for j := da.y; j <= a.y; j++ {
			luci[point{i, j}] = false
		}
	}
	return luci
}

func toggle(luci map[point]bool, da, a point) map[point]bool {
	for i := da.x; i <= a.x; i++ {
		for j := da.y; j <= a.y; j++ {
			if luci[point{i, j}] {
				luci[point{i, j}] = false
			} else {
				luci[point{i, j}] = true
			}
		}
	}
	return luci
}

func stampatrue(luci map[point]bool) int {
	var c int
	for _, v:= range luci{
		if v{
			c++
		}
	}
	return c
}
