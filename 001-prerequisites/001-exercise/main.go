package main

import "fmt"

type square struct {
	sisi float64
}

type circle struct {
	r float64
}

func (s square) sumArea() float64 {
	return s.sisi * s.sisi
}

func (c circle) sumArea() float64 {
	return 3.14 * (c.r * c.r)
}

type shape interface {
	sumArea() float64
}

func info(s shape) {
	fmt.Println("Area :", s.sumArea())
}

func main() {
	sq1 := square{5}
	ci1 := circle{7}

	info(sq1)
	info(ci1)
}
