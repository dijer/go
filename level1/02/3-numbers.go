package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Введите трехначное число:\n")
	fmt.Scan(&number)

	units := number % 10
	tens := int((number % 100) / 10)
	hundreds := int((number % 1000) / 100)

	fmt.Print(hundreds, " сотен, ", tens, " десятков, ", units, " единиц\n")
}
