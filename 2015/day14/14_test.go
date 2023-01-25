package main

import (
	"bufio"
	"fmt"
	"testing"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

func TestParteA(t *testing.T) {

	var giocatori []giocatore

	sc := bufio.NewScanner(zairo.Input("prova.txt"))

	var tempGiocatore giocatore
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &tempGiocatore.nome, &tempGiocatore.velocita, &tempGiocatore.tempo, &tempGiocatore.pausa)

		giocatori = append(giocatori, tempGiocatore)
	}

	risultato := parteA(giocatori, 1000)

	if risultato != 1120 {
		t.Errorf("A FAIL. Risultato: %d, ricevuto: %d", 1120, risultato)
	} else {
		t.Logf("A PASSED.")
	}

}
func TestParteB(t *testing.T) {

	var giocatori []giocatore

	sc := bufio.NewScanner(zairo.Input("prova.txt"))

	var tempGiocatore giocatore
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &tempGiocatore.nome, &tempGiocatore.velocita, &tempGiocatore.tempo, &tempGiocatore.pausa)

		giocatori = append(giocatori, tempGiocatore)
	}

	risultato := parteB(giocatori, 1000)

	if risultato != 689 {
		t.Errorf("B FAIL. Risultato: %d, ricevuto: %d", 689, risultato)
	} else {
		t.Logf("B PASSED.")
	}

}
