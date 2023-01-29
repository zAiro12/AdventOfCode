package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	file, _ := os.Open(os.Args[1])
	sc := bufio.NewScanner(file)
	defer file.Close()

	var contenitori []int
	var num int
	for sc.Scan() {
		num, _ = strconv.Atoi(sc.Text())
		contenitori = append(contenitori, num)
	}

	zairo.StampaA(parteA(contenitori, 25))
	zairo.StampaB(parteB(contenitori))
}

type elemento struct {
	val int
	pos int
}

func parteA(contenitori []int, totLitri int) int {

	var counter int

	sort.Sort(sort.Reverse(sort.IntSlice(contenitori)))

	visitati := make(map[elemento]bool)

	for i := 0; i < len(contenitori); i++ {
		visitati[elemento{contenitori[i], i}] = true

        zairo.Log("",counter, visitati, i)
		counter += cerca(contenitori, visitati, i, 0, contenitori[i], totLitri)

	}

	return counter
}

func cerca(contenitori []int, visitati map[elemento]bool, indice, daDove, tempCount, totLitri int) int {
	for i := daDove; i < len(contenitori); i++ {

		if !visitati[elemento{contenitori[i], daDove + i}] && tempCount+contenitori[i] <= totLitri {
			tempCount += contenitori[i]
		}

		if tempCount == totLitri {
			return 1 + cerca(contenitori, visitati, indice, i, contenitori[i], totLitri)
		}
	}
	return 0
}

func parteB(contenitori []int) int {

	return 0
}
