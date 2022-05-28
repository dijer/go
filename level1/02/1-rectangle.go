package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Введите a, b:\n")
	fmt.Scanln(&a, &b)

	area := a * b

	fmt.Print("Площадь прямоугольника: ", area, "\n")
}
