package main

import (
	"fmt"
	"strings"
)

var c int

func main() {

	for {
		var in string
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		split := strings.Split(in, "")
		ric(split)
	}

	fmt.Println(c)
}

func togli(slice []string, s int) []string {
	return append(slice[:s-1], slice[s+1:]...)
}

func ric(split []string) int {
a:
	for i := 0; i < len(split); i++ {
		switch split[i] {
		case ")":
			if split[i-1] == "(" {
				split = togli(split, i)
				return ric(split)
			} else {
				c += 3
				break a
			}

		case "]":
			if split[i-1] == "[" {
				split = togli(split, i)
				return ric(split)
			} else {
				c += 57
				break a
			}

		case "}":
			if split[i-1] == "{" {
				split = togli(split, i)
				return ric(split)
			} else {
				c += 1197
				break a
			}

		case ">":
			if split[i-1] == "<" {
				split = togli(split, i)
				return ric(split)
			} else {
				c += 25137
				break a
			}
		}
	}

	return c
}
