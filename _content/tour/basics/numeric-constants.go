//go:build OMIT

package main

import "fmt"

const (
	// Создаём огромное число, сдвигая бит 1 влево на 100 позиций.
	// Другими словами, двоичное число, которое представляет собой 1, за которой следуют 100 нулей.
	Big = 1 << 100
	// Сдвигаем его обратно вправо на 99 позиций, так что получаем 1<<1, или 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
