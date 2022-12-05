package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	var in string
	stack1 := make(map[int][]string)
	stack2 := make(map[int][]string)

	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		in = sc.Text()
		nPile := (len(in) + 1) / 4
		var quanti, da, a int

		if in != "" && in[0] != 'm' {
			index := 1
			for j := 1; j <= nPile; j++ {
				val := in[(index)]
				if unicode.IsLetter(rune(val)) {
					stack1[j] = append(stack1[j], string(val))
					stack2[j] = append(stack2[j], string(val))
				}
				index += 4
			}

		} else if in != "" && in[0] == 'm' {
			fmt.Sscanf(sc.Text(), "move %d from %d to %d", &quanti, &da, &a)

			stack1 = parteA(quanti, da, a, stack1)
			stack2 = parteB(quanti, da, a, stack2)

		}
	}

	var risultatoA, risultatoB string
	for i := 1; i <= len(stack1); i++ {
		risultatoA += stack1[i][0]
	}
	for i := 1; i <= len(stack2); i++ {
		risultatoB += stack2[i][0]
	}
	zairo.StampaA(risultatoA)
	zairo.StampaB(risultatoB)
}

func parteA(quanti, da, a int, stack map[int][]string) map[int][]string {
	var appoggio []string

	appoggio = append(appoggio, stack[da][:quanti]...)
	stack[da] = stack[da][quanti:]

	appoggio = zairo.ReverseSlice(appoggio)

	appoggio = append(appoggio, stack[a]...)
	stack[a] = appoggio

	return stack
}

func parteB(quanti, da, a int, stack map[int][]string) map[int][]string {
	var appoggio []string

	appoggio = append(appoggio, stack[da][:quanti]...)
	stack[da] = stack[da][quanti:]

	appoggio = append(appoggio, stack[a]...)
	stack[a] = appoggio

	return stack
}
