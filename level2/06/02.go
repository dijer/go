package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	runtime.GOMAXPROCS(2)
	runtime.Gosched()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	time.Sleep(time.Second)
}
