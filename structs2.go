package main

import "fmt"

type point struct {
	x, y int
}

type circle struct {
	point
	radius int
}

type wheel struct {
	circle
	spokes int
}

func main() {
	p := point{}
	p.x = 21
	p.y = 12

	fmt.Println(p.x, p.y)

	c := circle{}
	c.x = 21
	c.y = 12
	c.radius = 7
	fmt.Println(c.x, c.y, c.radius)

	w := wheel{}
	w.x = 21
	w.y = 12
	w.radius = 7
	w.spokes = 10
	fmt.Println(w.x, w.y, w.radius, w.spokes)
}
