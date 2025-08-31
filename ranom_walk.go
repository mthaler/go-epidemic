package main

import "math/rand/v2"

type position struct {
	x float64
	y float64
}

func newPosition(x, y float64) *position {
	return &position{x: x, y: y}
}

func step(p position) position {
	x := p.x
	y := p.y
	if rand.Float64() >= 0.5 {
		x += 1.0
	} else {
		x -= 1.0
	}
	if rand.Float64() >= 0.5 {
		x += 1.0
	} else {
		x -= 1.0
	}
	return position{x: x, y: y}
}
