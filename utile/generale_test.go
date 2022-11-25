package utile

import "testing"

func TestMax(t *testing.T) {

	risultato := Max(1, 3)

	if risultato != 3 {
		t.Errorf("Max(1,3) FAIL. Risultato: %d, ricevuto: %d", 3, risultato)
	} else {
		t.Logf("Max(1,3) PASSED.")
	}
}
func TestMin(t *testing.T) {

	risultato := Min(1, 3)

	if risultato != 1 {
		t.Errorf("Max(1,3) FAIL. Risultato: %d, ricevuto: %d", 1, risultato)
	} else {
		t.Logf("Max(1,3) PASSED.")
	}
}

func TestRimuoviElemento(t *testing.T) {

	aspettato := []int{1, 2, 3, 5}
	risultato := RimuoviElemento([]int{1, 2, 3, 4, 5}, 3)
	for i := 0; i < len(risultato); i++ {
		if aspettato[i] != risultato[i] {
			t.Errorf("RimuoviElemento(%d, 3) FAIL. Risultato: %d, ricevuto %d", risultato, risultato[i], aspettato[i])
		}
	}
	t.Logf("RimuoviElemento(%d, 3) PASSED.", aspettato)

}
