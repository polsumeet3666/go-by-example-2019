package main

import (
	"fmt"
	"math"
)

//interface
type geometry interface {
	area() float64
	perim() float64
}

// struct rect
type rect struct {
	width, height float64
}

// struct circle
type circle struct {
	radius float64
}

// interface methods implementation on rect
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//interface methods implementation on circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// generic method invokation
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
