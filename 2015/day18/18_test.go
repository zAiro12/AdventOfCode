package main

import (
	"bufio"
	"strings"
	"testing"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func TestParteA(t *testing.T) {

	sc := bufio.NewScanner(zairo.Input("prova.txt"))

	griglia := griglia{make(map[point]bool), 0, 0}
	var split []string

	for sc.Scan() {
		split = strings.Split(sc.Text(), "")

		for x := 0; x < len(split); x++ {
			if split[x] == "#" {
				griglia.griglia[point{griglia.maxY, x}] = true
			}
		}
		griglia.maxY++
	}
	griglia.maxX = len(split)

	risultato := ParteA(griglia, 4)
	aspettato := 4

	if risultato != aspettato {
		t.Errorf("A FAIL. Risultato: %d, ricevuto: %d", aspettato, risultato)
	} else {
		t.Logf("A PASSED.")
	}
}

func TestParteB(t *testing.T) {

	sc := bufio.NewScanner(zairo.Input("prova.txt"))

	griglia := griglia{make(map[point]bool), 0, 0}
	var split []string

	for sc.Scan() {
		split = strings.Split(sc.Text(), "")

		for x := 0; x < len(split); x++ {
			if split[x] == "#" {
				griglia.griglia[point{griglia.maxY, x}] = true
			}
		}
		griglia.maxY++
	}
	griglia.maxX = len(split)

	griglia.griglia[point{0, 0}] = true
	griglia.griglia[point{griglia.maxX - 1, 0}] = true
	griglia.griglia[point{0, griglia.maxY - 1}] = true
	griglia.griglia[point{griglia.maxX - 1, griglia.maxY - 1}] = true

	risultato := ParteB(griglia, 5)
	aspettato := 17

	if risultato != aspettato {
		t.Errorf("B FAIL. Risultato: %d, ricevuto: %d", aspettato, risultato)
	} else {
		t.Logf("B PASSED.")
	}
}
