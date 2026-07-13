package main

import "fmt"

type Point struct {
	x, y int
}

func setPoint(p *Point) {
	p.x = 67
	p.y = 68
}

func main() {
	var p Point
	setPoint(&p)
	fmt.Printf("Point{ x: %d, y: %d }\n", p.x, p.y)

}
