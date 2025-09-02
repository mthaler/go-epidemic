package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// World represents the game state.
type World struct {
	randomWalks [][]position
	number      int
}

// NewWorld creates a new world.
func NewWorld(number int) *World {
	w := &World{
		randomWalks: make([][]position, number),
		number:      number,
	}
	return w
}

// Update game state by one step.
func (w *World) Update() {
	for i := 0; i < w.number; i++ {
		positions := w.randomWalks[i]
		s := len(positions)
		lastPos := positions[s-1]
		newPosition := step(lastPos)
		positions = append(positions, newPosition)
		w.randomWalks[i] = positions
	}
}

// Draw current game state.
func (w *World) Draw(pix []byte) {
	for _, r := range w.randomWalks {
		for _, p := range r {
			// fill pixels with black
			for i := 0; i < len(pix); i++ {
				pix[i] = 0
			}

			x := int(p.x)
			y := int(p.y)
			// draw random walk
			pix[x+y*screenWidth] = 0xff
		}
	}
}

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
	world  *World
	pixels []byte
}

func (g *Game) Update() error {
	g.world.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, screenWidth*screenHeight*4)
	}
	g.world.Draw(g.pixels)
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{
		world: NewWorld(screenWidth, screenHeight, int((screenWidth*screenHeight)/10)),
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Game of Life (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
