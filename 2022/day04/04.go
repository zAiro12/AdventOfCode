package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	var in string
	var c1, c2 int

	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		in = sc.Text()
		c1 += parteA(in)
		c2 += parteB(in)
		//zairo.Log(in, c1)
	}

	zairo.StampaA(c1)
	zairo.StampaB(c2)
}

func parteA(s string) int {
	zone := strings.Split(s, ",")

	zona1 := strings.Split(zone[0], "-")
	zona2 := strings.Split(zone[1], "-")

	num11, _ := strconv.Atoi(zona1[0])
	num12, _ := strconv.Atoi(zona1[1])
	num21, _ := strconv.Atoi(zona2[0])
	num22, _ := strconv.Atoi(zona2[1])

	mappa1 := make(map[int]bool)
	mappa2 := make(map[int]bool)

	for i := num11; i <= num12; i++ {
		mappa1[i] = true
	}
	for i := num21; i <= num22; i++ {
		mappa2[i] = true
	}
	if len(mappa1) < len(mappa2) {
		for k := range mappa1 {
			if !mappa2[k] {
				return 0
			}
		}
		return 1
	} else {
		for k := range mappa2 {
			if !mappa1[k] {
				return 0
			}
		}
		return 1
	}
}

func parteB(s string) int {
	zone := strings.Split(s, ",")

	zona1 := strings.Split(zone[0], "-")
	zona2 := strings.Split(zone[1], "-")

	num11, _ := strconv.Atoi(zona1[0])
	num12, _ := strconv.Atoi(zona1[1])
	num21, _ := strconv.Atoi(zona2[0])
	num22, _ := strconv.Atoi(zona2[1])

	mappa1 := make(map[int]bool)
	mappa2 := make(map[int]bool)

	for i := num11; i <= num12; i++ {
		mappa1[i] = true
	}
	for i := num21; i <= num22; i++ {
		mappa2[i] = true
	}

	for k := range mappa1 {
		if mappa2[k] {
			return 1
		}
	}
	for k := range mappa2 {
		if mappa1[k] {
			return 1
		}
	}
	return 0
}
