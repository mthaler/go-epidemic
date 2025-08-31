package main

import "math/rand/v2"

type position struct {
	x float64
	y float64
}

func generateRandomWalk2D(x, y float64, steps int) []position {
	var positions []position
	positions = make([]position, 0)
	stepX := 0.0
	if rand.Float64() >= 0.5 {
		stepX = 1.0
	} else {
		stepX = -1.0
	}
	stepY := 0.0
	if rand.Float64() >= 0.5 {
		stepY = 1.0
	} else {
		stepY = -1.0
	}
	xx := x + stepX
	yy := y + stepY
	p := position{x: xx, y: yy}
	positions = append(positions, p)
	return positions
}
