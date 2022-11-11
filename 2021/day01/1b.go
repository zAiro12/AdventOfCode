package main

import "fmt"

func main() {
	var c, prev, curr int
	var slice []int
	for {
		_, err := fmt.Scan(&prev)
		if err != nil {
			break
		}
		slice = append(slice, prev)
	}
	prev = slice[0] + slice[1] + slice[2]
	for i := 3; i < len(slice); i++ {
		curr = prev - slice[i-3] + slice[i]
		if curr > prev {
			c++
		}
	}
	fmt.Println(c)
}
