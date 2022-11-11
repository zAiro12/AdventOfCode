package main

import "fmt"

func main() {
	var pesci []int
	for {
		var in int
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		pesci = append(pesci, in)
	}
	for i := 0; i < 80; i++ {
		for j := 0; i < len(pesci); j++ {
			if pesci[j] == 0 {
				pesci[j] = 7
				pesci = append(pesci, 9)
			}
			pesci[j]--
		}
	}
	fmt.Println(len(pesci))
}
