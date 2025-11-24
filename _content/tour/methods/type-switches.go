//go:build OMIT

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Удвоенное %v равно %v\n", v, v*2)
	case string:
		fmt.Printf("%q имеет длину %v байт\n", v, len(v))
	default:
		fmt.Printf("Я не знаю тип %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
