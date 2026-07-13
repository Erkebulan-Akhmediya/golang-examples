package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error")
		return
	}
	res := string(bytes[:len(bytes)-1])
	fmt.Println(res + "a")
}
