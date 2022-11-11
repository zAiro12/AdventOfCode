package main

import (
	"fmt"
)

func main() {
	var c, prev, curr int
	fmt.Scan(&curr)
	for {
		_, err := fmt.Scan(&curr)
		if err != nil {
			break
		}
		if curr > prev {
			c++
		}
		prev = curr
	}
	fmt.Println("uscita", c)
}
