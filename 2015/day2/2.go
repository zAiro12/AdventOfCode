package main

import (
	"fmt"
	"sort"
)

func main() {
	var somm int
	var s []int
	var in int
	for {
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		s = append(s, in)
	}
	for len(s) != 0 {
		//sub := s[0:4]
		l := s[0]
		w := s[1]
		h := s[2]
		fmt.Println(s[0], s[1], s[2])
		somm += (2 * l * w) + (2 * w * h) + (2 * h * l)
		sort.Ints(s[0:4])
		somm += s[1] * s[2]
		s = s[4:]
	}
	fmt.Println(somm)
}
