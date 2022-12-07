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

- NewFoglia : alloca uno nuovo spazio per una nuova foglia dell'abero
- GenAlbero : dato un array genera un albero completo o semicompleto
- StanmpaPreorder : stampa un albero in preorder
- StanmpaInorder : stampa un albero in inorder
- StanmpaPostorder : stampa un albero in postorder
- StampaAlberoASommario : stampa un albero a sommario
- StampaAlbero : stampa un albero

## GRAFI
> ### **INT**

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

> ### STRING

- **STRUCT**

    ```GOLANG
        type GrafoString map[string][]string
    ```

- NewGrafoString() : alloca un nuovo spazio per la mappa del grafo

- *InputString()* : **METODO** per leggere in input gl'archi del grafo

- *StampaGrafoString()* : **METODO** per stampare il grafo

## STRING

- IsLower : data una stringa retituisce se è lowercase
- IsUpper : data una stringa retituisce se è uppercase
- IsDiverse : data uan stringa restituisce se le lettere sono tutte diverse

## SLICE

- AggiungiElemento : dato un indice, aggiunge un elemento da una slice
- RimuoviElemento : dato un indice, rimuove un elemento da una slice
- AggiungiElementoTesta : aggiunge un elemento in prima posizione della slice
- AggiungiElementoCoda : aggiunge un elemento in ultima posizione della slice
- RimuoviElementoTesta : rimuovi un elemento in prima posizione della slice
- RimuoviElementoCoda : rimuovi un elemento in ultima posizione della slice
- SommaSlice : data una slice di int retituisce la somam degli elementi all'interno
- ReverseSlice : data una slice di qualsiasi tipo, retituisce la slice girata
- RimuoviVal : cerca un elemento e lo rimuove in tutte le posizioni dalla slice

## MATH

- Max : dati due numeri, restituisce il massimo
- Min : dati due numeri, restituisce il minimo

## RESTO

- **STRUCT**

    ```GOLANG
        type Point struct {
            X, Y int
        }
    ```

- Log : funzione per stampa di debug

- StampaA : funzione per stampare risultato della parte A

- StampaB : funzione per stampare risultato della parte B
