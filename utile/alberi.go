package utile

type NodoAlbero struct {
	Destra   *NodoAlbero
	Sinistra *NodoAlbero
	Val      int
}

type Albero struct {
	Root *NodoAlbero
}

func NewFoglia(val int) *NodoAlbero {
	return &NodoAlbero{nil, nil, val}
}

func Aggiungi(val int, Albero *Albero) {
	foglia := NewFoglia(val)
	appoggio := Albero.Root
	for appoggio.Sinistra != nil {
		appoggio = appoggio.Sinistra
	}
	for appoggio.Destra != nil {
		appoggio = appoggio.Destra
	}
	appoggio = foglia
}
