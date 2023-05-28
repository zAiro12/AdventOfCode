package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type operazione struct{
	campo1 string
	operazione string
	campo2 string
	not bool
}

func main(){

	valori := make(map[string]int)
	operazioni := make(map[operazione]string)
	
	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)

	var in string
	var split []string
	var split2 []string
	var num int
	var op operazione


	for sc.Scan(){
		in = sc.Text()
		split = strings.Split(in, " -> ")

		valori[split[1]] = 0

		split2 = strings.Split(split[0], " ")
		if (len(split2) == 1){

			num, _ = strconv.Atoi(split2[0])
			valori[split[1]] = num

		}else if (len(split2) == 2){
			op = operazione{split2[1], "NOT","", true}
			operazioni[op] = split[1]
		}else{
			op = operazione{split2[0], split2[1], split2[2], false}
			operazioni[op] = split[1]
		}
	}

	zairo.Logln("valori",len(valori), valori)
	zairo.Logln("operazioni", len(operazioni), operazioni)

	zairo.StampaA(parteA(valori, operazioni))
	zairo.StampaB(parteB())
}

func parteA(valori map[string]int, operazioni map[operazione]string) map[string]int{
	
}

func parteB() int{
	return 0
}