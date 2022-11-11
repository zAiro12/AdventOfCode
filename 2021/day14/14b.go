package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var pol string
	fmt.Scan(&pol)
	finale := make(map[string]uint64)
	for i := 0; i < len(pol)-1; i++ {
		chiave := fmt.Sprintf("%c%c", pol[i], pol[i+1])
		finale[chiave]++
	}

	operazioni := make(map[string]string)
	var uno, due string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Text()
		fmt.Sscanf(in, "%s -> %s", &uno, &due)
		if uno == "" {
			continue
		}
		operazioni[uno] = due
	}

	for i := 0; i < 40; i++ {
		finale = passo(finale, operazioni)
	}

	ricorenze := make(map[rune]uint64)
	ricorenze[rune(pol[0])]++
	for k, v := range finale {
		ricorenze[rune(k[1])] += v
	}

	var max, min uint64 = 0, ricorenze[rune(pol[0])]
	for k, v := range ricorenze {
		if ricorenze[k] > max {
			max = v
		} else if ricorenze[k] < min {
			min = v
		}
	}
	fmt.Println(max - min)
}

func passo(finale map[string]uint64, operazioni map[string]string) map[string]uint64 {
	nuova := make(map[string]uint64)
	var chiave1, chiave2 string
	for k, v := range finale {
		chiave1 = string(k[0]) + operazioni[k]
		chiave2 = operazioni[k] + string(k[1])
		nuova[chiave1] += v
		nuova[chiave2] += v
	}
	return nuova
}
