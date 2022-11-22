package utile

type Nodo struct {
	destra   *Nodo
	sinistra *Nodo
	val      int
}

type Albero struct {
	root *Nodo
}

func NewFoglia(val int) *Nodo {
	return &Nodo{nil, nil, val}
}

func Aggiungi(val int, Albero *Albero) {
	foglia := NewFoglia(val)
	appoggio := Albero.root
	for appoggio.sinistra != nil {
		appoggio = appoggio.sinistra
	}
	for appoggio.destra != nil {
		appoggio = appoggio.destra
	}
	appoggio = foglia
}
