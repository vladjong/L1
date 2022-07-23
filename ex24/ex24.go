package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p *Point) Init(x int, y int) {
	p.x = x
	p.y = y
}

func newPoint(x int, y int) *Point {
	point := Point{x, y}
	return &point
}

func (p *Point) DistanceBetweenPoints() int {
	if p.x == p.y {
		return 0
	} else if p.x > p.y {
		return p.x - p.y
	} else {
		return p.y - p.x
	}
}

func main() {
	point := newPoint(10, 4)
	dis := point.DistanceBetweenPoints()
	fmt.Println(dis)
}
