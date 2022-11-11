package main

import (
	"fmt"
	"math"
)

func main() {
	var input, oxigen, carbon []string
	var n string
	mapp := make(map[int]int)
	// read the input
	for {
		fmt.Scan(&n)
		if n == "end" {
			break
		}
		input = append(input, n)
	}
	// count the first column
	for i := 0; i < len(input); i++ {
		mapp[int(input[i][0])-48]++
	}
	//fmt.Println(mapp)
	// First divider
	if mapp[0] > mapp[1] {
		for i := 0; i < len(input); i++ {
			if input[i][0] == '0' {
				oxigen = append(oxigen, input[i])
			} else {
				carbon = append(carbon, input[i])
			}
		}
	} else {
		for i := 0; i < len(input); i++ {
			if input[i][0] == '1' {
				oxigen = append(oxigen, input[i])
			} else {
				carbon = append(carbon, input[i])
			}
		}
	}
	//fmt.Println(len(oxigen), len(carbon))
	for i := 1; i < 12; i++ {
		oxigen = CalcO(oxigen, i)
	}
	for i := 1; i < 12; i++ {
		carbon = CalcC(carbon, i)
	}
	// conversion
	var ox, car float64
	for i := 0; i < 12; i++ {
		ox += float64(oxigen[0][i]-48) * math.Pow(2, float64(11-i))
		car += float64(carbon[0][i]-48) * math.Pow(2, float64(11-i))
	}
	fmt.Println(ox * car)
}

func CalcO(input []string, n int) []string {
	if len(input) == 1 {
		return input
	}
	mapp := make(map[int]int)
	var oxigen []string
	for i := 0; i < len(input); i++ {
		mapp[int(input[i][n])-48]++
	}
	//fmt.Println(mapp)
	// First divider
	if mapp[0] > mapp[1] {
		for i := 0; i < len(input); i++ {
			if input[i][n] == '0' {
				oxigen = append(oxigen, input[i])
			}
		}
	} else {
		for i := 0; i < len(input); i++ {
			if input[i][n] == '1' {
				oxigen = append(oxigen, input[i])
			}
		}
	}
	return oxigen
}

func CalcC(input []string, n int) []string {
	if len(input) == 1 {
		return input
	}
	mapp := make(map[int]int)
	var carbon []string
	for i := 0; i < len(input); i++ {
		mapp[int(input[i][n])-48]++
	}
	//fmt.Println(mapp)
	// First divider
	if mapp[0] > mapp[1] {
		for i := 0; i < len(input); i++ {
			if input[i][n] == '1' {
				carbon = append(carbon, input[i])
			}
		}
	} else {
		for i := 0; i < len(input); i++ {
			if input[i][n] == '0' {
				carbon = append(carbon, input[i])
			}
		}
	}
	return carbon
}
