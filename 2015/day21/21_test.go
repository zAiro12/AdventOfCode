package main

import (
	"fmt"
	"testing"
)

func TestProva(t *testing.T) {
	giocatore := entita{8, 5, 5}
	boss := entita{12, 7, 2}

	var i int
	for {

		if boss.vita <= 0 || giocatore.vita <= 0 {
			break
		}

		if i%2 == 0 {
			if giocatore.danno-boss.armatura > 0 {

				boss.vita -= giocatore.danno - boss.armatura
			} else {
				boss.vita--
			}
		} else {
			if boss.danno-giocatore.armatura > 0 {

				giocatore.vita -= boss.danno - giocatore.armatura
			} else {
				giocatore.vita--
			}
		}

		i++
	}

	fmt.Println(giocatore.vita, boss.vita)
	if giocatore.vita > boss.vita && boss.vita <= 0 {
		t.Logf("A PASSED.")
	} else {
		t.Errorf("A FAIL.")
	}
}
