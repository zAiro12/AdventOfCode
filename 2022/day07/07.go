package main

import (
	"bufio"
	"fmt"
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

	stampaAlbero(cartelle.root)
	zairo.StampaA(parteA())
	zairo.StampaB()
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

		} else if split[1] == "cd" && split[2] != "/" && split[2] != ".." {
			cursore = setCursore(split[2], cursore)
			zairo.Log("FUORI cursore:", cursore)
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
	return node
}

func aggiungiSottoCartella(s string, cartellaCorrente *cartella) {
	node := newNode(s, cartellaCorrente)
	node.sottocartelle = append(node.sottocartelle, node)
	zairo.Log("figli creati", node, node.sopra)
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
func parteA() int {

	return 0
}
