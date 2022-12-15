package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type cartelle struct {
	root *cartella
}

type cartella struct {
	nome          string
	dimensione    int
	sopra         *cartella
	sottocartelle []*cartella
}

func main() {
	var cartelle cartelle
	creaAlbero(&cartelle)

	aggiornavalori(cartelle.root)
	zairo.StampaA(parteA(cartelle))
	zairo.StampaB(parteB(cartelle))
}

func creaAlbero(cartelle *cartelle) {
	file := zairo.Input(os.Args[1])
	sc := bufio.NewScanner(file)

	var in string
	var split []string

	cursore := creaRoot("/")
	cartelle.root = cursore

	for sc.Scan() {
		in = sc.Text()
		split = strings.Split(in, " ")

		if split[1] == "cd" && split[2] == ".." {
			cursore = cursore.sopra

		} else if split[1] == "cd" && split[2] != ".." {
			cursore = setCursore(split[2], cursore)
		}

		if split[1] == "ls" {
			for sc.Scan() {
				in2 := sc.Text()
				split2 := strings.Split(in2, " ")

				if split2[0] == "dir" {
					aggiungiSottoCartella(split2[1], cursore)

				} else if split2[0] == "$" {

					if split2[2] == ".." {
						cursore = cursore.sopra
					} else {
						cursore = setCursore(split2[2], cursore)
					}
					break

				} else {
					num, _ := strconv.Atoi(split2[0])
					cursore.dimensione += num
				}

			}
		}
	}
}

func newNode(s string, father *cartella) *cartella {
	return &cartella{s, 0, father, make([]*cartella, 0)}
}

func creaRoot(s string) *cartella {
	node := newNode(s, nil)
	node.sopra = node
	return node
}

func aggiungiSottoCartella(s string, cartellaCorrente *cartella) {
	node := newNode(s, cartellaCorrente)
	cartellaCorrente.sottocartelle = append(cartellaCorrente.sottocartelle, node)
}

func setCursore(s string, cursore *cartella) *cartella {
	for _, v := range cursore.sottocartelle {
		if v.nome == s {
			return v
		}
	}
	return cursore
}

func stampaAlbero(cartella *cartella) {

	if cartella == nil {
		return
	}
	fmt.Println(cartella.nome, cartella.dimensione, cartella.sottocartelle)
	for i := 0; i < len(cartella.sottocartelle); i++ {
		stampaAlbero(cartella.sottocartelle[i])
	}
}

func aggiornavalori(cartella *cartella) {
	for i := 0; i < len(cartella.sottocartelle); i++ {
		aggiornavalori(cartella.sottocartelle[i])
		cartella.dimensione += cartella.sottocartelle[i].dimensione
	}
}

func parteA(cartelle cartelle) int {
	var daVisitare []*cartella
	var counter int

	daVisitare = append(daVisitare, cartelle.root)
	var cursore *cartella

	for len(daVisitare) != 0 {
		cursore = daVisitare[0]
		daVisitare = daVisitare[1:]

		if cursore.dimensione < 100000 {
			counter += cursore.dimensione
		}

		if len(cursore.sottocartelle) != 0 {
			daVisitare = append(daVisitare, cursore.sottocartelle...)
		}
	}
	return counter
}

func parteB(dir cartelle) int {

	libero := 70000000 - dir.root.dimensione
	daCercare := 30000000 - libero

	var frontiera []*cartella
	frontiera = append(frontiera, dir.root)
	min := math.MaxInt

	for len(frontiera) != 0 {
		if frontiera[0].dimensione > daCercare && frontiera[0].dimensione < min {
			min = frontiera[0].dimensione
		} else {
			frontiera = append(frontiera, frontiera[0].sottocartelle...)
			frontiera = frontiera[1:]
		}
	}
	return min
}
