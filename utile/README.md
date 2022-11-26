# PACKAGE CREATO DA ME

## PILA

- NewNodePila : alloca uno nuovo spazio per un nuovo nodo
- PushPila : aggiunge un nuovo elemento in testa alla pila
- PopPila : toglie l'elemento in testa alla pila
- StampaPila : stampa la pila

## ALBERI

- **STRUCT** (NON ANCORA FINITO)

    ```GOLANG
    type NodoAlbero struct {
	    Destra   *NodoAlbero
	    Sinistra *NodoAlbero
	    Val      int
    }

    type Albero struct {
	    Root *NodoAlbero
    }
    ```

- NewFoglia : alloca uno nuovo spazio per una nuovoa foglia
- Aggiungi : aggiunge una nuova foglia all'albero

## GRAFI

- **STRUCT**

    ```GOLANG
        type NodopilaInt struct {
            Val  int
            Next *Nodopila
        }

        type GrafoInt struct {
            Len       int
            Adiacenti []*Nodopila
        }
    ```

- NuovoGrafoInt : alloca lo spazio per un nuovo grafo

- AggiungiCoppieGrafoInt : dato un vertice (x), crea un arco con un altro vertice (y)

- StampaGrafoInt : stampa il grafo vertice per vertice

- IsArcoGrafoInt : dice se esiste un arco tra due nodi

## RESTO

- **STRUCT**

    ```GOLANG
        type Point struct {
            X, Y int
        }
    ```

- RimuoviElemento : dato un indice, rimuove un elemento da una slice
- Max : dati due numeri, restituisce il massimo
- Min : dati due numeri, restituisce il minimo
