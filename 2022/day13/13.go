package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {

	sc := bufio.NewScanner(zairo.Input(os.Args[1]))

	var in string
	for sc.Scan() {
		in = sc.Text()
		riga1 := riempiRiga(in)
		stampaSlice(riga1)

		sc.Scan()
		in = sc.Text()
		riga2 := riempiRiga(in)
		stampaSlice(riga2)

		sc.Scan()
	}

	zairo.StampaA()
	zairo.StampaB()
}

func riempiRiga(s string) (input []any) {
	s = s[1 : len(s)-1]
	split := strings.Split(s, ",")
	for i := 0; i < len(split); i++ {
		input = append(input, split[i])
	}
	return
}

func stampaSlice(slice []any) {
	for i := 0; i < len(slice); i++ {
		zairo.Log("slice", i, slice[i])
	}
	fmt.Println()
}
