package main

import (
	"fmt"
)

func main() {
	var c [9]int
	var in int
	for {
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		c[in]++
	}
	for i := 1; i <= 256; i++ {
		c[0], c[1], c[2], c[3], c[4], c[5], c[6], c[7], c[8] = c[1], c[2], c[3], c[4], c[5], c[6], c[7]+c[0], c[8], c[0]

	}

	fmt.Println(c[0] + c[1] + c[2] + c[3] + c[4] + c[5] + c[6] + c[7] + c[8])
}
