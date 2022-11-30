package main

import (
	"fmt"
	"strings"
)

func main() {
	var in string
	mappa := make(map[string]string)
	for {
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		trim := strings.Split(in, ")")
		mappa[trim[1]] = trim[0]
	}

	Stampa(mappa)

	var c int
	for k := range mappa {
		c += conta(k, mappa)
	}
	fmt.Println(c)

}

func conta(s string, m map[string]string) int {
	if len(m[s]) != 0 {
		return 1 + conta(m[s], m)
	} else {
		return 0
	}
}

func Stampa(mappa map[string]string) {
	for k, v := range mappa {
		fmt.Println(k, ")", v)
	}
}
