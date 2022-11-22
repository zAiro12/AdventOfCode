package utile

type NodoAlbero struct {
	destra   *NodoAlbero
	sinistra *NodoAlbero
	val      int
}

type Albero struct {
	root *NodoAlbero
}

func NewFoglia(val int) *NodoAlbero {
	return &NodoAlbero{nil, nil, val}
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
