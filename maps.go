package main

import "fmt"

func main() {
	ages := map[string]int{
		"bob":  21,
		"ivan": 30,
		"john": 0,
	}
	ages["sara"] = 34
	ages["bob"] = 12
	for name, age := range ages {
		fmt.Println(name, age)
	}
	age, ok := ages["john"]
	if !ok {
		fmt.Println("john is not present")
	}
	fmt.Println("john", age)
}
