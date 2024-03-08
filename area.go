package main

import "fmt"

type shape interface {
	getArea() float64
}

type traingle struct {
	height float64
	base   float64
}

type squire struct {
	sideLenth float64
}

func (t traingle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s squire) getArea() float64 {
	return s.sideLenth * s.sideLenth
}

func printArea(s shape) {
	fmt.Println("Area :", s.getArea())
}
