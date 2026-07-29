package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := filter(nums, func(n int) bool {
		return n > 5
	})
	fmt.Println(res)
}

// функция пробегает по срезу nums вызывая f для важдого элемента
// и записывает элемент при котором f возвращает true в новый срез
// и возвращает его
func filter(nums []int, f func(int) bool) []int {
	var res []int
	for _, num := range nums {
		if f(num) {
			res = append(res, num)
		}
	}
	return res
}
