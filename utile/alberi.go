package utile

type nodo struct {
	destra   *nodo
	sinistra *nodo
	val      int
}

type albero struct {
	root *nodo
}

func newFoglia(val int) *nodo {
	return &nodo{nil, nil, val}
}

func aggiungi(val int, albero *albero) {
	foglia := newFoglia(val)
	appoggio := albero.root
	for appoggio.sinistra != nil {
		appoggio = appoggio.sinistra
	}
	for appoggio.destra != nil {
		appoggio = appoggio.destra
	}
	appoggio = foglia
}
