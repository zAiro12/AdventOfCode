package main

import "fmt"

func main() {

	fmt.Println(parteA())
	fmt.Println(parteB())

}

func parteA() int {
	const k = 34000000 / 11
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

	return 0
}

func parteB() int {
	const k = 34000000 / 11
	var i int
	casa := make(map[int][]int)
	elfi := make(map[int]int)
	var som int
	for i = 1; ; i++ {
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				if elfi[j] < 50 {
					casa[i] = append(casa[i], j)
					elfi[j]++
				}
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
	fmt.Println("ris", i)

	return 0
}
