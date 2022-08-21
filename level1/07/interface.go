package main

import (
	"fmt"
	"io"
)

type MyReader struct{}

func (reader MyReader) Write(p []byte) (int, error) {
	fmt.Printf("%s + from MyReader Write\n", string(p))
	return len(p), nil
}

func (reader MyReader) Read(p []byte) (int, error) {
	fmt.Printf("%s + from MyReader Read\n", p)
	return len(p), nil
}

func main() {
	io.WriteString(MyReader{}, "test")
	io.ReadAll(MyReader{})
}
