package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}

	fmt.Println("arr == arr2", arr == arr2)
	fmt.Printf("%T\n", arr)
	for i, n := range arr {
		fmt.Println(i, n)
	}
}
