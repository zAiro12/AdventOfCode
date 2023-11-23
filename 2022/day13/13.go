package main

import (
	"bufio"
	"os"
	"strconv"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	sc := bufio.NewScanner(zairo.Input(os.Args[1]))

	var lettura1, lettura2 string
	var index int
	for sc.Scan() {
		index++
		lettura1 = sc.Text()
		sc.Scan()
		lettura2 = sc.Text()

		lettura1 = lettura1[1 : len(lettura1)-1]
		lettura2 = lettura2[1 : len(lettura2)-1]

		res, err := parteA(lettura1, lettura2)
		zairo.Log("index", index, res, err)

		sc.Scan()
	}

	zairo.StampaA()
	zairo.StampaB()
}

func parteA(str1, str2 string) (bool, int) {
	len1 := len(str1)
	len2 := len(str2)

	if len1 == 0 && len2 != 0 {
		return true, 0
	}

	if len2 == 0 && len1 != 0 {
		return false, 0
	}

	var i1, i2 int
	var ris string
	for i1 < len1 && i2 < len2 {
		ris = comparaElemento(str1[i1], str2[i2], str1, str2)

		switch ris {
		case "=":
			i1 += 2
			i2 += 2
			break

		case "<":
			return true, 0

		case ">":
			return true, 0

		case "1[":
			zairo.Log("1[")
		case "2[":
			zairo.Log("2[")

		default:
			zairo.Log("switch", ris, string(str1[i1]), string(str2[i2]))
		}
	}

	return false, 1
}

func isNum(s byte) bool {
	_, err := strconv.Atoi(string(s))

	if err != nil {
		return false
	}
	return true
}

func comparaElemento(el1, el2 byte, str1, str2 string) string {
	if isNum(el1) && isNum(el2) {
		num1, _ := strconv.Atoi(string(el1))
		num2, _ := strconv.Atoi(string(el2))

		if num1 < num2 {
			return "<"
		}

		if num2 < num1 {
			return ">"
		}

		return "="

	}
	if el1 == '[' && el2 == '['{


		parteA(str1, str2)
		return "["
	}

	return "err"
}
