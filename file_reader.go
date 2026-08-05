package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return // os.Exit(1)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bytes))

}
