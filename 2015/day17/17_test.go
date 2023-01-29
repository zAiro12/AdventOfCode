package main

import "testing"

func TestParetA(t *testing.T) {

	contenitori := []int{20, 15, 10, 5, 5}

	risultato := parteA(contenitori, 25)
	aspettato := 4
	if risultato != aspettato {
		t.Errorf("A FAIL. Risultato: %d, ricevuto: %d", aspettato, risultato)
	} else {
		t.Logf("A PASSED.")
	}

}

func TestParetB(t *testing.T) {

	contenitori := []int{20, 15, 10, 5, 5}

	risultato := parteB(contenitori)
	aspettato := 0
	if risultato != aspettato {
		t.Errorf("B FAIL. Risultato: %d, ricevuto: %d", aspettato, risultato)
	} else {
		t.Logf("B PASSED.")
	}

}
