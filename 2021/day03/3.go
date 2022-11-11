package main

import (
	"fmt"
	"strconv"
)

func main() {
	var in, gammaBIN, epBIN string
	var s []string
	map1 := make(map[int]int)
	map0 := make(map[int]int)
	for {
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		s = append(s, in)
	}
	for i := 0; i < 12; i++ {
		for j := 0; j < len(s); j++ {
			if s[j][i] == '1' {
				map1[i]++
			} else {
				map0[i]++
			}
		}
	}
	for i := 0; i < 12; i++ {
		if map0[i] > map1[i] {
			gammaBIN += "0"
			epBIN += "1"
		} else {
			gammaBIN += "1"
			epBIN += "0"
		}
	}
	gamma, _ := strconv.ParseInt(gammaBIN, 2, 64)
	ep, _ := strconv.ParseInt(epBIN, 2, 64)

	fmt.Println(gamma * ep)

}
