package main

import (
	"fmt"
)

func main() {
	n := abc(1, 2, 3)
	fmt.Println(n)
	a, _, c := f(1) // если некоторые значения не нужны, их можно заменить на _
	fmt.Println(a, c)
	fmt.Println(sum(2, 2))
}

func abc(a, b, c int) int {
	return 0
}

func f(a int) (int, int, int) {
	return 1, 2, 3
}

// еще один способ как можно вернуть значение из фукнции
func sum(a, b int) (c int) {
	c = a + b
	return
}
