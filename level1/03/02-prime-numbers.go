package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите N:")
	fmt.Scanln(&n)

	var prime []int

	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			prime = append(prime, i)
		}
	}

	fmt.Println("Простые числа: ", prime)
}
