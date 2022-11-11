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

	stampa(mappa)

	fmt.Println(trova(mappa, "SAN"))

}

func trova(m map[string]string, s string) int {
	return 0
}

func stampa(mappa map[string]string) {
	for k, v := range mappa {
		fmt.Println(k, ")", v)
	}
}
