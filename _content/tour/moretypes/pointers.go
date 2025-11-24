//go:build OMIT

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // указываем на i
	fmt.Println(*p) // читаем i через указатель
	*p = 21         // устанавливаем i через указатель
	fmt.Println(i)  // видим новое значение i

	p = &j         // указываем на j
	*p = *p / 37   // делим j через указатель
	fmt.Println(j) // видим новое значение j
}
