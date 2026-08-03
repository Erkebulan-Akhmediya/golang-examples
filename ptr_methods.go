package main

import "fmt"

type Point struct {
	x, y int
}

func (p *Point) Scale(n int) {
	p.x *= n
	p.y *= n
}

// func (p Point) Scale(n int) {}

func main() {
	p := Point{2, 3}
	p.Scale(2)
	fmt.Println(p.x, p.y) // 4 6
}
