package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, !): ")
	fmt.Scanln(&op)
	if op != "!" {
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
	}
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
			os.Exit(1)
		}
		res = a / b
	case "^":
		res = float64(math.Pow(a, b))
	case "!":
		res = 1
		for i := 1; i <= int(a); i++ {
			res *= float64(i)
		}
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}
