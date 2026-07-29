package main

import "fmt"

func main() {
	generator := createGenerator()
	fmt.Println(generator())
	fmt.Println(generator())
	fmt.Println(generator())
}

// возвращает нам функцию,
// каждый вызов которой возвращает значение на 1 больше предыдущего
func createGenerator() func() int {
	var n int
	return func() int {
		n++
		return n
	}
}
