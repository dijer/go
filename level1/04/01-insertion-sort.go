package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	var slice []int
	fmt.Println("Вводите числа для входящего массива, нажимая enter.\nКогда закончите - нажмите enter на пустой строке")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		} else {

			number, err := strconv.Atoi((scanner.Text()))
			if err == nil {
				slice = append(slice, number)
			}
		}
	}

	fmt.Println("Входящий массив:\n", slice)

	insertionSort(slice)

	fmt.Println("\nОтсортированный массив:\n", slice)
}
