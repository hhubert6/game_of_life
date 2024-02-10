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
	WIDTH      = 640
	HEIGHT     = 360
	rectSize   = 1
	gridWidth  = WIDTH / rectSize
	gridHeight = HEIGHT / rectSize
)

var tempGrid [][]int8

type Game struct {
	grid [][]int8
}

func makeGrid() [][]int8 {
	grid := make([][]int8, gridHeight)
	for i := range grid {
		grid[i] = make([]int8, gridWidth)
	}

	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			if rand.Float32() < 0.4 {
				grid[y][x] = int8(1)
			}
		}
	}
	return grid
}

func init() {
	tempGrid = make([][]int8, gridHeight)
	for i := range tempGrid {
		tempGrid[i] = make([]int8, gridWidth)
	}
}

func (g *Game) Update() error {
	for y := range g.grid {
		for x := range g.grid[y] {
			tempGrid[y][x] = g.rule(x, y)
		}
	}

	tmp := g.grid
	g.grid = tempGrid
	tempGrid = tmp

	return nil
}

func (g *Game) rule(x, y int) int8 {
	cnt := int8(0)
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			target_x, target_y := x+dx, y+dy

			if (target_x != x || target_y != y) &&
				target_x >= 0 && target_x < gridWidth &&
				target_y >= 0 && target_y < gridHeight {
				cnt += g.grid[target_y][target_x]
			}
		}
	}

	alive := g.grid[y][x] == 1
	if alive && (cnt == 2 || cnt == 3) || !alive && cnt == 3 {
		return int8(1)
	}
	return int8(0)
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%.2f", ebiten.ActualFPS()))

	for i := range g.grid {
		for j := range g.grid[i] {
			if g.grid[i][j] == 1 {
				drawRect(screen, j*rectSize, i*rectSize)
			}
		}
	}
}

func drawRect(screen *ebiten.Image, x, y int) {
	vector.DrawFilledRect(screen, float32(x), float32(y), rectSize, rectSize, color.White, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func main() {
	game := &Game{grid: makeGrid()}

	ebiten.SetTPS(30)
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Game of life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
