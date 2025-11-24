//go:build OMIT

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("Значение:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("Значение:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("Значение:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("Значение:", v, "Присутствует?", ok)
}
