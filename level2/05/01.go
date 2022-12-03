package main

import "fmt"

func main() {
	n := 1000
	ch := make(chan bool)

	for i := 1; i < 5; i++ {
		go func() {
			counter := 0
			for k := 1; k <= 5; k++ {
				counter++
				fmt.Println("go", n, "-", counter)
			}
			ch <- true
		}()
	}
	for i := 1; i < 5; i++ {
		<-ch
	}
	fmt.Println("The End")
}
