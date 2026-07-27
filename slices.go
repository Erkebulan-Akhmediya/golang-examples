package main

import "fmt"

func main() {
	s := make([]int, 0)
	for range 12 {
		s = append(s, 10)
		fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	}

	var n []int
	fmt.Println("n == nil", n == nil)

	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	fmt.Println("a == b", equals(a, b))
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
