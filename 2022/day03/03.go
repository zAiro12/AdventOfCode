package main

import (
	"bufio"
	"os"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func main() {
	zairo.StampaA(parteA())
	zairo.StampaB(parteB())
}

func parteA() int {
	var meta, counter int
	var in string

	input := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(input)

	for sc.Scan() {
		in = sc.Text()
		meta = len(in) / 2

		primaParte := in[:meta]
		secondaParte := in[meta:]

		lettere1 := mappaString(primaParte)
		lettere2 := mappaString(secondaParte)

		intruso := trovaIntruso(lettere1, lettere2)
		counter+= calcolaPunteggio(intruso)
	}
	return counter
}

func calcolaPunteggio(lettera byte) (counter int) {
	if zairo.IsLower(string(lettera)) {
		counter += int(lettera - 'a' + 1)
	} else {
		counter += int(lettera - 'A' + 27)
	}
	return
}

func mappaString(s string) map[byte]bool {
	lettere := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		lettere[s[i]] = true
	}
	return lettere
}

func trovaIntruso(prima, seconda map[byte]bool) byte {
	for k := range prima {
		if seconda[k] {
			return k
		}
	}
	return 0
}

func parteB() (c int) {

	input := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(input)

	for sc.Scan() {
		in1 := sc.Text()
		sc.Scan()
		in2 := sc.Text()
		sc.Scan()
		in3 := sc.Text()

		lettere1 := mappaString(in1)
		lettere2 := mappaString(in2)
		lettere3 := mappaString(in3)

		badge := trovaBadge(lettere1, lettere2, lettere3)
		c += calcolaPunteggio(badge)
	}
	return
}

func trovaBadge(prima, seconda, terza map[byte]bool) byte {
	var possibili []byte

	for k := range prima {
		if seconda[k] {
			possibili = append(possibili, k)
		}
	}

	for i := 0; i < len(possibili); i++ {
		if terza[possibili[i]] {
			return possibili[i]
		}
	}
	return 0
}
