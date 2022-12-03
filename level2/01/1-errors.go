package main

import (
	"errors"
	"fmt"
	"time"
)

func someFunc(a int, b int) (int, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if b == 0 {
		now := time.Now()

		return 0, errors.New(now.Format("01-02-2006 15:04:05"))
	}

	return a / b, nil
}

func main() {
	var div, err = someFunc(3, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("outputs", div)
}
