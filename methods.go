package main

import "fmt"

type IntSlice []int

func (s IntSlice) ForEach(f func(int)) {
	for _, n := range s {
		f(n)
	}
}

func (s IntSlice) ForEachMut(f func(*int)) {
	for i := range s {
		f(&s[i])
	}
}

func (s IntSlice) Map(f func(int) int) IntSlice {
	s2 := make(IntSlice, 0, len(s))
	for _, n := range s {
		s2 = append(s2, f(n))
	}
	return s2
}

func main() {
	slice := IntSlice{1, 2, 3, 4, 5}
	slice.ForEach(func(n int) {
		fmt.Println(n)
	})

	slice2 := slice.Map(func(n int) int {
		return n * 2
	})
	fmt.Println(slice2)

	slice.ForEachMut(func(n *int) {
		*n *= -1
	})
	fmt.Println(slice)
}
