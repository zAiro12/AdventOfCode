package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var pol string
	fmt.Scan(&pol)
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
		pol = step(pol, operazioni)
		fmt.Println(i, len(pol))
	}

	ricorenze := make(map[rune]int)
	for _, k := range pol {
		ricorenze[k]++
	}
	max, min := 0, len(pol)

	for k, v := range ricorenze {
		if ricorenze[k] > max {
			max = v
		} else if ricorenze[k] < min {
			min = v
			fmt.Println(min)
		}
	}
	fmt.Println(max-min)
}

func step(s string, operazioni map[string]string) string {
	var chiave string
	for i := 0; i < len(s)-1; i++ {
		chiave = fmt.Sprintf("%c%c", s[i], s[i+1])
		if operazioni[chiave] != "" {
			s = s[:i+1] + operazioni[chiave] + s[i+1:]
			i++
		}
	}
	return s
}
