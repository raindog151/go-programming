package main

import (
	"fmt"
	"math"
)

type square struct {
	length int
	width int
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main() {

	mySquare := square {
		length: 10,
		width: 80,
	}

	myCircle := circle {
		radius: 5.0,
	}

	info(mySquare)
	info(myCircle)

}

func (c circle) area() (float64) {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() (float64) {
	return float64(s.length * s.width)
}

func info(s shape) {
	fmt.Println(s.area())
}
