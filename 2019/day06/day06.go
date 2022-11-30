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

	fmt.Println(risolviA(padri))
	fmt.Println(risolviB(padri))

}

//risolvi parte a usando funzione conta
func risolviA(padri map[string]string) int {
	var c int
	for k := range padri {
		c += conta(k, padri)
	}
	return c
}

//parte a
func conta(s string, m map[string]string) int {
	if len(m[s]) != 0 {
		return 1 + conta(m[s], m)
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

	for i := 0; san != "COM" || you != "COM"; i++ {
		san = padri[san]
		you = padri[you]
		fmt.Println("debug:", san, you)

		visitatiSan[san] = i
		visitatiYou[you] = i

	}
	fmt.Println(visitatiSan)
	fmt.Println(visitatiYou)
	return passi
}


func stampa(mappa map[string]string) {
	for k, v := range mappa {
		fmt.Println(k, ")", v)
	}
}
