//go:build OMIT

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Моё любимое число -", rand.Intn(10))
}
