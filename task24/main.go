package main

import (
	"fmt"
	"math"
)


type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (this Point) Distance(other Point) float64 {
	return  math.Sqrt(math.Pow(this.x - other.x, 2) + math.Pow(this.y - other.y, 2))
}

func main() {
	this := NewPoint(0, 0)
	other := NewPoint(4, 4)
	fmt.Println(this.Distance(*other))
}