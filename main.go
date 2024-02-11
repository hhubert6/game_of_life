package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 360
	RECT_SIZE     = 1
	GRID_WIDTH    = SCREEN_WIDTH / RECT_SIZE
	GRID_HEIGHT   = SCREEN_HEIGHT / RECT_SIZE
)

type Grid [][]int8

type Game struct {
	Grid     Grid
	TempGrid Grid
}

func NewGame() *Game {
	game := &Game{NewGrid(), NewGrid()}

	for x := 0; x < SCREEN_WIDTH; x++ {
		for y := 0; y < SCREEN_HEIGHT; y++ {
			if rand.Float32() < 0.2 {
				game.Grid[y][x] = int8(1)
			}
		}
	}

	return game
}

func NewGrid() Grid {
	grid := make(Grid, GRID_HEIGHT)
	for i := range grid {
		grid[i] = make([]int8, GRID_WIDTH)
	}

	return grid
}

func (g *Game) Update() error {
	for y := range g.Grid {
		for x := range g.Grid[y] {
			g.TempGrid[y][x] = g.Rule(x, y)
		}
	}

	g.Grid, g.TempGrid = g.TempGrid, g.Grid

	return nil
}

func (g *Game) Rule(x, y int) int8 {
	cnt := int8(0)
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			target_x, target_y := x+dx, y+dy

			if (target_x != x || target_y != y) &&
				target_x >= 0 && target_x < GRID_WIDTH &&
				target_y >= 0 && target_y < GRID_HEIGHT {
				cnt += g.Grid[target_y][target_x]
			}
		}
	}

	alive := g.Grid[y][x] == 1
	if alive && (cnt == 2 || cnt == 3) || !alive && cnt == 3 {
		return int8(1)
	}
	return int8(0)
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%.2f", ebiten.ActualFPS()))

	for i := range g.Grid {
		for j := range g.Grid[i] {
			if g.Grid[i][j] == 1 {
				DrawRect(screen, j*RECT_SIZE, i*RECT_SIZE)
			}
		}
	}
}

func DrawRect(screen *ebiten.Image, x, y int) {
	vector.DrawFilledRect(screen, float32(x), float32(y), RECT_SIZE, RECT_SIZE, color.White, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	game := NewGame()

	ebiten.SetTPS(30)
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Game of life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
