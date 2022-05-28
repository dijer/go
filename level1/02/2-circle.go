package main

import (
	"fmt"
	"math"
)

func main() {
	var area float64
	fmt.Print("Введите площадь круга:\n")
	fmt.Scan(&area)

	diameter := 2 * math.Sqrt(area/math.Pi)
	fmt.Print("Диаметр окружности: ", math.Floor(diameter*100)/100, "\n")

	var circumference float64 = math.Pi * diameter
	fmt.Print("Длина окружности: ", math.Floor(circumference*100)/100, "\n")
}
