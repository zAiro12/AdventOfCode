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

	fmt.Println("A:", risolviA(padri))
	fmt.Println("B:", risolviB(padri))

}

// risolvi parte a usando funzione conta
func risolviA(padri map[string]string) int {
	var c int
	for k := range padri {
		c += conta(k, padri)
	}
	return c
}

// parte a
func conta(s string, padri map[string]string) int {
	if len(padri[s]) != 0 {
		return 1 + conta(padri[s], padri)
	} else {
		return 0
	}
}

func risolviB(padri map[string]string) int {
	var passi int
	visitatiSan := make(map[string]int)
	visitatiYou := make(map[string]int)
	san := "SAN"
	you := "YOU"

	for i := 0; ; i++ {
		if san == "COM" || you == "COM" {
			break
		}
		san = padri[san]
		you = padri[you]

		visitatiSan[san] = i
		visitatiYou[you] = i

		if visitatiSan[you] > 0 {
			passi = visitatiSan[you] + visitatiYou[you]
			break
		}
		if visitatiYou[san] > 0 {
			passi = visitatiYou[san] + visitatiSan[san]
			break
		}
	}
	return passi
}

func stampa(mappa map[string]string) {
	for k, v := range mappa {
		fmt.Println(k, ")", v)
	}
}
