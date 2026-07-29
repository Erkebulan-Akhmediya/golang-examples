package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5} // 15

	// создание функции внутри функции main
	print := func(n int) {
		fmt.Println(n)
	}

	forEach(nums, print) //

	var sum int
	calcSum := func(n int) {
		sum += n
	}
	forEach(nums, calcSum)
	fmt.Println(sum)
}

// функция пробегается в срезу nums вызывая f для каждого эленмета
func forEach(nums []int, f func(int)) {
	for _, num := range nums {
		f(num)
	}
}
