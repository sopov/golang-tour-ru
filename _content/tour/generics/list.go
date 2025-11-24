//go:build OMIT

package main

// List представляет односвязный список, который содержит
// значения любого типа.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
}
