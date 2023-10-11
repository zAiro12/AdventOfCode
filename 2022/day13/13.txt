package main

import (
	"bufio"
	"os"
	"strconv"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type pacchetto struct {
	numeri map[int]int
	lista  map[int]pacchetto
}

func main() {

	sc := bufio.NewScanner(zairo.Input(os.Args[1]))

	var in string
	var counterA, index int
	for sc.Scan() {
		in = sc.Text()
		pacchetto1 := riempiRiga(in)

		sc.Scan()
		in = sc.Text()
		pacchetto2 := riempiRiga(in)

		sc.Scan()
		index++

		//zairo.Logln("pachhetti", pacchetto1, "\n", pacchetto2)
		if parteA(pacchetto1, pacchetto2) < 0 {
			counterA += index
		} else {
			zairo.Logln("index", index)
		}

	}

	zairo.StampaA(counterA)
	zairo.StampaB()
}

func riempiRiga(s string) pacchetto {
	p := pacchetto{make(map[int]int), make(map[int]pacchetto)}

	var index, contaQuadre int
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '[' {
			contaQuadre++

			for j := i; j < len(s); j++ {

				if s[j] == ']' {
					contaQuadre--

					if contaQuadre == 0 {
						p.lista[index] = riempiRiga(s[i : j+1])
						index++
					}

					i = j + 1
					break
				}
			}

		} else {
			for j := i; j < len(s); j++ {
				if s[j] == ',' || j == len(s)-1 {
					n, err := strconv.Atoi(s[i:j])

					if err == nil {
						if n == 0 {
							n = -1
						}
						p.numeri[index] = n
						index++
					}
					i = j
					break
				}
			}

		}
	}
	return p
}

func parteA(pacchetto1, pacchetto2 pacchetto) int {
	var tmp int
	if pacchetto1.isEmpty() && pacchetto2.isEmpty() {
		return 0
	}
	if pacchetto1.isEmpty() {
		return -1
	}
	if pacchetto2.isEmpty() {
		return 1
	}

	lunghezza := zairo.Max(len(pacchetto1.lista)+len(pacchetto1.numeri), len(pacchetto2.lista)+len(pacchetto2.numeri))

	for i := 0; i < lunghezza; i++ {

		zairo.Log("condizioni", pacchetto1, pacchetto2, pacchetto1.numeri[i] != 0, pacchetto2.numeri[i] != 0)
		if pacchetto1.numeri[i] != 0 && pacchetto2.numeri[i] != 0 {
			if pacchetto1.numeri[i] < pacchetto2.numeri[i] {
				return -1
			}
			if pacchetto1.numeri[i] > pacchetto2.numeri[i] {
				return 1
			}
			continue
		}

		if pacchetto1.numeri[i] != 0 {
			return parteA(convert(pacchetto1.numeri[i]), pacchetto2.lista[i])
		}
		if pacchetto2.numeri[i] != 0 {
			return parteA(pacchetto1.lista[i], convert(pacchetto2.numeri[i]))
		}

		tmp = parteA(pacchetto1.lista[i], pacchetto2.lista[i])

		if tmp != 0 {
			return tmp
		}

		continue
	}

	return 0
}

func (p pacchetto) isEmpty() bool {
	return len(p.lista)+len(p.numeri) == 0
}

// Returns a Packet made using a number
func convert(num int) pacchetto {
	return riempiRiga("[" + strconv.Itoa(num) + "]")
}
