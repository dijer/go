package main

import "fmt"

var cache map[int]int = make(map[int]int, 0)

func fibonacci(n int) int {
	if cache[n] != 0 {
		return cache[n]
	}
	if n <= 1 {
		return n
	}
	result := fibonacci(n-1) + fibonacci(n-2)
	cache[n] = result
	return result
}

func main() {
	var n int
	fmt.Println("Введите n-е число Фибоначчи:")
	fmt.Scanln(&n)

	result := fibonacci(n)
	fmt.Println("Число Фибоначчи:", result)
}
