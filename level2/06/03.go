package main

func main() {
	b := 1000

	go func() {
		b++
	}()

	b--
}
