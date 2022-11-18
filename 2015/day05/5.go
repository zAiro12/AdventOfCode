package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var in string
	var c int
	for sc.Scan() {
		in = sc.Text()
		if isNiceB(in) {
			c++
		}
	}
	fmt.Println(c)
}

// parte B
func isNiceB(s string) bool {
	sillabe := make(map[string]int)
	var sillaba string
	check := false
	for i := 0; i < len(s)-2; i++ {
		sillaba = fmt.Sprintf("%c%c", s[i], s[i+1])
		if s[i] == s[i+1] {
			check = true
		}
		sillabe[sillaba]++
	}
	sillaba = fmt.Sprintf("%c%c", s[len(s)-2], s[len(s)-1])
	sillabe[sillaba]++

	for _, v := range sillabe {
		if v >= 2 && check {
			return true
		}
	}
	return false
}

// parte A
func isNice(s string) bool {
	lettere := make(map[byte]int)
	sillabe := make(map[string]int)
	var sillaba string
	for i := 0; i < len(s)-1; i++ {
		lettere[s[i]]++
		sillaba = fmt.Sprintf("%c%c", s[i], s[i+1])
		sillabe[sillaba]++
	}
	lettere[s[len(s)-1]]++

	if sillabe["ab"] > 0 || sillabe["cd"] > 0 || sillabe["pq"] > 0 || sillabe["xy"] > 0 {
		// fmt.Println("Sillaba vietata:", s)
		return false
	}
	if vocali(lettere) == false {
		// fmt.Println("Vocali mancanti:", s)
		return false
	}
	if doppie(sillabe) == false {
		// fmt.Println("Sillaba doppia mancante:", s)
		return false
	}

	return true
}

func vocali(mappa map[byte]int) bool {
	lung := mappa['a'] + mappa['e'] + mappa['i'] + mappa['o'] + mappa['u']
	return lung >= 3
}

func doppie(mappa map[string]int) bool {
	for k := range mappa {
		if k[0] == k[1] {
			return true
		}
	}
	return false
}

func stampa(mappa map[byte]int) {
	for k, v := range mappa {
		fmt.Printf("%c , %d   ", k, v)
	}
	fmt.Println()
}
