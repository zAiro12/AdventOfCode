package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	zairo.StampaA(parteA())
	zairo.StampaB(parteB())
}

func parteA() int {
	var in string
	var current, max, n int

	input := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(input)

	for sc.Scan() {
		in = sc.Text()

		if in == "" {
			if current > max {
				max = current
			}
			current = 0
		} else {
			n, _ = strconv.Atoi(in)
			current += n
		}
	}
	return max
}

func parteB() int {
	var in string
	var current, n int
	max := []int{0, 0, 0}

	input := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(input)

	for sc.Scan() {
		in = sc.Text()
		if in == "" {
			max = aggiornaMax(max, current)
			current = 0
		} else {
			n, _ = strconv.Atoi(in)
			current += n
		}
	}
	return zairo.SommaSlice(max)
}

func aggiornaMax(max []int, val int) []int {
	max = append(max, val)
	sort.Ints(max)
	return max[1:]
}
