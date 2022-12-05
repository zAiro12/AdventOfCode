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
	stack1 := make(map[int][]string)
	stack2 := make(map[int][]string)

	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)
	for i := 0; i < 10; i++ {
		sc.Scan()

		in = sc.Text()
		split := strings.Split(in, "")
		if i == 9 {
			break
		}
		stack1[i+1] = append(stack1[i+1], split...)
		stack2[i+1] = append(stack2[i+1], split...)
	}

	var quanti, da, a int
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &quanti, &da, &a)

		stack1 = parteA(quanti, da, a, stack1)
		stack2 = parteB(quanti, da, a, stack2)
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
