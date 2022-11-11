package main

import (
	"fmt"
	"math"
)

func main() {
	var num []int
	for {
		var in int
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		num = append(num, in)
	}
	var min int = 234567890
	for col := 0; col < 2000; col++ {
		carbu := 0
		for i := 0; i < len(num); i++ {
			carbu += step(int(math.Abs(float64(num[i]) - float64(col))))
			if carbu > min {
				break
			}
		}
		if carbu < min {
			min = carbu
		}
	}
	fmt.Println(min)
}

func step(n int) int {
	var out int
	for i := 0; i <= n; i++ {
		out += i
	}
	return out
}
