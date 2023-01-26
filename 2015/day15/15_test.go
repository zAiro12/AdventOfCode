package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestParteA(t *testing.T) {

	biscotti := make(map[int]biscotto)

	file, _ := os.Open("prova.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)

	var appoggio biscotto
	var i int
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &appoggio.nome, &appoggio.capacita, &appoggio.durabilita, &appoggio.sapore, &appoggio.texture, &appoggio.calorie)

		appoggio.nome = strings.Trim(appoggio.nome, ":")
		biscotti[i] = appoggio
		i++
	}

	risultato := parteA(biscotti)

	if risultato != 62842880 {
		t.Errorf("B FAIL. Risultato: %d, ricevuto: %d", 62842880, risultato)
	} else {
		t.Logf("B PASSED.")
	}
}
func TestParteB(t *testing.T) {

	biscotti := make(map[int]biscotto)

	file, _ := os.Open("prova.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)

	var appoggio biscotto
	var i int
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &appoggio.nome, &appoggio.capacita, &appoggio.durabilita, &appoggio.sapore, &appoggio.texture, &appoggio.calorie)

		appoggio.nome = strings.Trim(appoggio.nome, ":")
		biscotti[i] = appoggio
		i++
	}

	risultato := parteB(biscotti)

	if risultato != 0 {
		t.Errorf("B FAIL. Risultato: %d, ricevuto: %d", 689, risultato)
	} else {
		t.Logf("B PASSED.")
	}
}
