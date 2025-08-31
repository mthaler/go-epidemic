package main

type position struct {
	x float64
	y float64
}

func generateRandomWalk2D(x, y float64, steps int) []position {
	var positions []position
	positions = make([]position, 0)
	xx := x
	yy := y
	p := position{x: xx, y: yy}
	positions = append(positions, p)
	return positions
}
