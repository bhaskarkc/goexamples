// Package mystructs provides ...
package mystructs

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Printf("Passed type %T has an area: %v\n", s, s.area())
}

func Methodset() {
	c := circle{5}
	info(&c)
}
