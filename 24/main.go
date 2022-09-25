package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) Init(x, y float64) {
	p.x = x
	p.y = y
}

func Dist(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	p1 := new(Point)
	p2 := new(Point)
	p1.Init(2, 10)
	p2.Init(15, 23)
	fmt.Println(Dist(p1, p2))
}
