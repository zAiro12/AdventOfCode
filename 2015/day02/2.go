package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var somm int
	var in []string
	var input string
	var l, w, h int

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input = sc.Text()
		in = strings.Split(input, "x")
		l, _ = strconv.Atoi(in[0])
		w, _ = strconv.Atoi(in[1])
		h, _ = strconv.Atoi(in[2])
		somm += day2(l, w, h)
	}

	fmt.Println(somm)
}

func day1(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l + min(l, w, h)
}

func min(x, y, z int) int {
	appoggio := []int{x, y, z}
	sort.Ints(appoggio)
	return appoggio[1] * appoggio[2]
}


func day2(l, w, h int) int {
	appoggio := []int{l, w, h}
	sort.Ints(appoggio)
	return min2(l, w, h) + l*w*h
}

func min2(x, y, z int) int {
	appoggio := []int{x, y, z}
	sort.Ints(appoggio)
	return 2*appoggio[0] + 2*appoggio[1]
}
