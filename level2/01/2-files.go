package main

import (
	"fmt"
	"os"
	"strconv"
)

func createFiles() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("an error occured: %v", e)
		}
	}()

	for i := 0; i < 1000000; i++ {
		file, err := os.Create(strconv.Itoa(i) + ".txt")
		defer file.Close()
		file.Close()
		if err != nil {
			return err
		}
	}
	return
}

func main() {
	err := createFiles()
	if err != nil {
		fmt.Println(err)
	}
}
