package utile

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type Integer interface {
	Signed | Unsigned
}
type Float interface {
	~float32 | ~float64
}
type Ordered interface {
	Integer | Float
}

// data una slice di qualsiasi tipo aggiunge l'elemento in posizione index
func AggiungiElemento[T comparable](arr []T, index int, val T) []T {
	a := arr[:index]
	a = append(a, val)
	a = append(a, arr[index:]...)
	return a
}

// data una slice di qualsiasi tipo elimina l'elemento in posizione index
func RimuoviElemento[T comparable](arr []T, index int) []T {
	a := arr[:index]
	a = append(a, arr[index+1:]...)
	return a
}

// aggiunge un elemento in prima posizione della slice
func AggiungiElementoTesta[T comparable](arr []T, elemento T) []T {
	var a []T
	a = append(a, elemento)
	a = append(a, arr...)
	return a
}

// aggiunge un elemento in ultima posizione della slice
func AggiungiElementoCoda[T comparable](arr []T, elemento T) []T {
	arr = append(arr, elemento)
	return arr
}

// rimuovi un elemento in prima posizione della slice
func RimuoviElementoTesta[T comparable](arr []T) []T {
	return arr[1:]
}

// rimuovi un elemento in ultima posizione della slice
func RimuoviElementoCoda[T comparable](arr []T) []T {
	return arr[:len(arr)-1]
}

// data una slice di int retituisce la somma degli elementi all'interno
func SommaSlice(arr []int) int {
	var c int
	for i := 0; i < len(arr); i++ {
		c += arr[i]
	}
	return c
}

// data una slice di qualsiasi tipo, retituisce la slice girata
func ReverseSlice[T comparable](arr []T) []T {
	lunghezza := len(arr)
	for i := 0; i < lunghezza/2; i++ {
		arr[i], arr[lunghezza-1-i] = arr[lunghezza-1-i], arr[i]
	}

	return arr
}

// cerca un elemento e lo rimuove in tutte le posizioni dalla slice
func RimuoviVal[T comparable](arr []T, elemento T) []T {
	for i, v := range arr {
		if v == elemento {
			arr = RimuoviElemento(arr, i)
		}
	}
	return arr
}

// data una slice di int retituisce la moltiplicazione degli elementi all'interno
func MoltipicaSlice[T Ordered](arr []T) T {
	var c T
	c = 1
	for i := 0; i < len(arr); i++ {
		c *= arr[i]
	}
	return c
}
