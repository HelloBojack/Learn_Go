package main

import (
	"fmt"
	"math"
)

const PI = math.Pi

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * PI
}

func (c circle) perim() float64 {
	return 2 * c.radius * PI
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{10, 20}
	c := circle{10}

	measure(r)
	measure(c)
}
