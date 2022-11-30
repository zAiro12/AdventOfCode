package main

import (
	"fmt"
	"strings"
)

func main() {
	var in string
	padri := make(map[string]string)
	for {
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		trim := strings.Split(in, ")")
		padri[trim[1]] = trim[0]
	}
	stampa(padri)

	//fmt.Println(distanza(padri))

}

func distanza(padri map[string]string) int {
	var passi int
	visitatiSan := make(map[string]int)
	visitatiYou := make(map[string]int)
	san := "SAN"
	you := "YOU"

	for i := 0; san == "COM" || you == "COM"; i++ {
		san = padri[san]
		you = padri[you]

		visitatiSan[san] = i
		visitatiYou[you] = i
		
	}

	return passi
}

func stampa(mappa map[string]string) {
	for k, v := range mappa {
		fmt.Println(k, ")", v)
	}
}
