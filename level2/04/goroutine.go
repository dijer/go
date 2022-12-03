package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	workers := make(chan int, 5)

	for i := 0; i < 1000; i++ {
		workers <- 1

		go func(job int) {
			defer func() {
				count += <-workers
			}()
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("total value", count)
}
