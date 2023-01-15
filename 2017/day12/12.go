package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Compilare gli archi in base all'input
	grafo := createGraph()
	// Effettuare ricerca (BFS?) utilizzando la EQL
	// Restituire lunghezza di EQL
	fmt.Println("Parte A:", len(BFS(grafo, "0")))
	fmt.Println("Parte B:", quantiGruppi(grafo))
}

// Crea il grafo leggendo da input

func createGraph() map[string][]string {
	grafo := make(map[string][]string)

	file, _:= os.Open(os.Args[1])
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		riga := strings.Split(sc.Text(), " <-> ")
		grafo[riga[0]] = strings.Split(riga[1], ", ")
	}
	return grafo
}

// Effettua ricerca in ampiezza del grafo e restituisce quanti nodi ha visiatato
func BFS(graph map[string][]string, start string) map[string]bool {
	visitati := make(map[string]bool)
	coda := make([]string, 0)
	
	visitati[start] = true
	coda = append(coda, start)
	
	for len(coda) > 0 {
		// Scegli nodo da espandere
		nodoDaAnalizzare := coda[0]
		coda = coda[1:]
		// Espandi nodo
		for i := 0; i < len(graph[nodoDaAnalizzare]); i++ {
			// Aggiungo a F
			if !visitati[graph[nodoDaAnalizzare][i]] {
				coda = append(coda, graph[nodoDaAnalizzare][i])
				// Aggiungo a EQL i nodi visitati
				visitati[graph[nodoDaAnalizzare][i]] = true
			}
		}
	}
	return visitati
}

// Restituisce quanti gruppi esistono (Quante volte viene chiamata BFS)
func quantiGruppi(grafo map[string][]string) int {
	counter := 0
	// Chiama il primo BFS
	visitati := BFS(grafo, "0")
	counter++
	// Ripeto per ogni nodo non ancora visitato
	for i := 0; i < len(grafo); i++ {
		tmp := strconv.Itoa(i)
		if !visitati[tmp] {
			temp := BFS(grafo, tmp)
			for k := range temp {
				visitati[k] = true
			}
			counter++
		}
	}
	return counter
}