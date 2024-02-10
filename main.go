package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	WIDTH  = 640
	HEIGHT = 360
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func main() {
	game := &Game{}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Game of life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
