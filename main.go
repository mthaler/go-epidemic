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

// Update game state by one tick.
func (w *World) Update() {
}

// Draw paints current game state.
func (w *World) Draw(pix []byte) {
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
