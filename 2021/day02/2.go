package main

import (
	"fmt"
	"strconv"
)

func main() {
	var mem []string
	var in string
	var x, y int
	for {
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		mem = append(mem, in)
	}

	for i := 0; i < len(mem); i += 2 {
		if mem[i] == "forward" {
			x1, _ := strconv.Atoi(mem[i+1])
			x += x1
		} else if mem[i] == "down" {
			y1, _ := strconv.Atoi(mem[i+1])
			y += y1
		} else if mem[i] == "up" {
			y1, _ := strconv.Atoi(mem[i+1])
			y -= y1
		}
	}
	fmt.Println(mem[0])
	fmt.Println(x * y)
}
