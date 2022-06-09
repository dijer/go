package main

import "fmt"

func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		j := i
		for j > 0 {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
			j = j - 1
		}
	}
}

func main() {
	var slice []int = []int{4, 101, 100, 5, 1, 3, 1}
	insertionSort(slice)

	fmt.Println("Отсортированный массив:")
	fmt.Println(slice)
}
