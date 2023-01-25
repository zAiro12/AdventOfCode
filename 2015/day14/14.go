package main

import (
	"bufio"
	"fmt"
	"os"

	zairo "github.com/zAiro12/AdventOfCode/utile"
)

type giocatore struct {
	nome     string
	velocita int
	tempo    int
	pausa    int
}

func main() {
	var giocatori []giocatore

	sc := bufio.NewScanner(zairo.Input(os.Args[1]))

	var tempGiocatore giocatore
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &tempGiocatore.nome, &tempGiocatore.velocita, &tempGiocatore.tempo, &tempGiocatore.pausa)

		giocatori = append(giocatori, tempGiocatore)
	}

	zairo.StampaA(parteA(giocatori, 2503))
	zairo.StampaB(parteB(giocatori, 2503))
}

func parteA(giocatori []giocatore, SECONDI_TOTALI int) int {
	var max int

	for i := 0; i < len(giocatori); i++ {
		durataAzione := giocatori[i].tempo + giocatori[i].pausa
		spostamenti := SECONDI_TOTALI / durataAzione
		rimSec := SECONDI_TOTALI % durataAzione

		tempSpazio := spostamenti * giocatori[i].velocita * giocatori[i].tempo

		if rimSec <= giocatori[i].tempo {
			tempSpazio += rimSec * giocatori[i].velocita
		} else {
			tempSpazio += giocatori[i].velocita * giocatori[i].tempo
		}

		if tempSpazio > max {
			max = tempSpazio
		}
	}
	return max
}

func parteB(giocatori []giocatore, SECONDI_TOTALI int) int {

	type tipo struct {
		posCorrente int
		valFinale   int
	}

	counter := make(map[string]tipo, len(giocatori))

	for i := 0; i < SECONDI_TOTALI; i++ {
		for _, valoreGiocatore := range giocatori {

			// zairo.Log("MOD", i%(valoreGiocatore.pausa+valoreGiocatore.tempo))
			if i%(valoreGiocatore.pausa+valoreGiocatore.tempo) < valoreGiocatore.tempo {

				appoggio := counter[valoreGiocatore.nome]
				appoggio.posCorrente += valoreGiocatore.velocita
				counter[valoreGiocatore.nome] = appoggio

			}
		}

		var max int
		var nome string
		for _, valoreGiocatore := range giocatori {

			if counter[valoreGiocatore.nome].posCorrente > max {

				max = counter[valoreGiocatore.nome].posCorrente
				nome = valoreGiocatore.nome

			}
		}
		appoggio := counter[nome]
		appoggio.valFinale++
		counter[nome] = appoggio
		zairo.Logln("Mappa", i, counter)
	}

	var max int
	for _, v := range counter {
		if v.valFinale > max {
			max = v.valFinale
		}
	}
	return max
}
