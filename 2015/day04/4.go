package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	input := "iwrupvqb"

	zairo.StampaA(risolviA(input))
	zairo.StampaB(risolviB(input))
}

func risolviA(input string) int {
	var c int
	for {
		c++
		caso := input + strconv.Itoa(c)

		cyp := md5.New()
		io.WriteString(cyp, caso)

		conversione := fmt.Sprintf("%x \n", cyp.Sum(nil))

		if conversione[:5] == "00000" {
			return c
		}
	}
}
func risolviB(input string) int {
	var c int
	for {
		c++
		caso := input + strconv.Itoa(c)

		cyp := md5.New()
		io.WriteString(cyp, caso)

		conversione := fmt.Sprintf("%x \n", cyp.Sum(nil))

		if conversione[:6] == "000000" {
			return c
		}
	}
}
