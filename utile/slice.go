package utile

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

func SommaSlice(arr []int) int {
	var c int
	for i := 0; i < len(arr); i++ {
		c += arr[i]
	}
	return c
}
