//go:build OMIT

package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Привет", World)
	fmt.Println("Счастливого", Pi, "дня")

	const Truth = true
	fmt.Println("Go рулит?", Truth)
}
