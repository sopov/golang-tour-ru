//go:build OMIT

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("привет", "мир")
	fmt.Println(a, b)
}
