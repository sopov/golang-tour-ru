//go:build OMIT

package main

import "fmt"

func main() {
	fmt.Println("считаем")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("готово")
}
