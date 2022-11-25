package utile

import "testing"

func TestMax(t *testing.T) {

	risultato := Max(1, 3)

	if risultato != 3{
		t.Errorf("Max(1,3) FAIL. Risultato: %d, ricevuto: %d", 3, risultato)
	}else{
		t.Logf("Max(1,3) PASSED.")
	}
}
