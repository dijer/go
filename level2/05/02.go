package main

import (
	"fmt"
	"sync"
)

const n = 1000

func main() {
	counter := 0
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(n)

	for i := 0; i < n; i += 1 {
		go func(j int, mutex *sync.Mutex) {
			mutex.Lock()
			defer mutex.Unlock()

			counter += 1

			fmt.Println("go", j+1)
			wg.Done()
		}(i, &mutex)
	}

	wg.Wait()

	fmt.Println("Counter: ", counter)
}
