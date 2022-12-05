package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	var in string
	stack := make(map[int][]string)

	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)
	for i := 0; i < 9; sc.Scan() {
		i++

		in = sc.Text()
		split := strings.Split(in, " ")
       zairo.Log(split)
		for i := 0; i < len(split); i++ {
			stack[i+1] = append(stack[i+1], split[i])
		}
		fmt.Println(stack[1], len(stack[1]))
	}

	zairo.StampaA()
	zairo.StampaB()
}

func parteA(s string) {

}
