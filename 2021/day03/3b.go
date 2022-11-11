package main

import (
	"fmt"
	"strconv"
)

func main() {
	var in string
	var s []string
	for {
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		s = append(s, in)
	}

	ossigeno := s
	carb := s
	for i := 0; i < 13; i++ {
		ossigeno = oss(ossigeno, i, true)
		carb = oss(carb, i, false)
		fmt.Println(ossigeno, carb, i)
	}
	fmt.Println(ossigeno, carb)
	ossy, _ := strconv.ParseInt(ossigeno[0], 2, 64)
	co2, _ := strconv.ParseInt(carb[0], 2, 64)

	fmt.Println(ossy * co2)
}

func oss(ossy []string, i int, isOssigeno bool) []string {
	var c1, c0 int
	var s1, s0 []string
	if len(ossy) == 1 {
		return ossy
	}
	for j := 0; j < len(ossy); j++ {
		if ossy[j][i] == '1' {
			s1 = append(s1, ossy[j])
			c1++
		}
		if ossy[j][i] == '0' {
			s0 = append(s0, ossy[j])
			c0++
		}
	}
	if isOssigeno {
		if c1 >= c0 {
			return s1
		} else {
			return s0
		}
	} else {
		if c0 >= c1 {
			return s0
		} else {
			return s1
		}
	}
}
