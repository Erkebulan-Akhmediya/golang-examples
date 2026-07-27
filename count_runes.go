package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	runeCount := make(map[rune]int)
	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		// EOF == end of file
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		runeCount[r]++
	}
	for r, count := range runeCount {
		fmt.Println(string(r), count)
	}
}
