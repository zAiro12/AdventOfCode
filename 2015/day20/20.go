package main

import "fmt"

func main() {
	const k = 34000000/11
	var i int
	casa := make(map[int][]int)
	var som int
	for i = 1; ; i++ {
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				casa[i] = append(casa[i], j)
			}

		}
		for _, v := range casa[i] {
			som += v
		}
		if som >= k {
			break
		}
		fmt.Println(som)
		som = 0
	}
	fmt.Println("ris", i, som)
}
