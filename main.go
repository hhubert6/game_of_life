package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	WIDTH      = 640
	HEIGHT     = 360
	rectSize   = 1
	gridWidth  = WIDTH / rectSize
	gridHeight = HEIGHT / rectSize
)

var rectImg *ebiten.Image
var grid [][]int8

type Game struct{}

func init() {
	rectImg = ebiten.NewImage(rectSize, rectSize)
	rectImg.Fill(color.RGBA{255, 255, 255, 255})

	grid = make([][]int8, gridHeight)
	for i := range grid {
		grid[i] = make([]int8, gridWidth)
	}

	for x := WIDTH / 4; x < 3*WIDTH/4; x++ {
		for y := HEIGHT / 3; y < 2*HEIGHT/3; y++ {
			if rand.Float32() < 0.2 {
				grid[y][x] = int8(1)
			}
		}
	}
}

func (g *Game) Update() error {
	tempGrid := make([][]int8, gridHeight)
	for i := range grid {
		tempGrid[i] = make([]int8, gridWidth)
	}

	for y := range grid {
		for x := range grid[y] {
			tempGrid[y][x] = rule(x, y)
		}
	}

	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = tempGrid[y][x]
		}
	}

	return nil
}

func rule(x, y int) int8 {
	cnt := int8(0)
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			target_x, target_y := x+dx, y+dy
			if (target_x != x || target_y != y) &&
				target_x >= 0 && target_x < gridWidth &&
				target_y >= 0 && target_y < gridHeight {
				cnt += grid[target_y][target_x]
			}
		}
	}

	alive := grid[y][x] == 1
	if alive && (cnt == 2 || cnt == 3) || !alive && cnt == 3 {
		return int8(1)
	}
	return int8(0)
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%.2f", ebiten.ActualFPS()))

	for i := range grid {
		for j := range grid[i] {
			if alive := grid[i][j]; alive == 1 {
				drawRect(screen, j*rectSize, i*rectSize)
			}
		}
	}
}

func drawRect(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(rectImg, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func main() {
	game := &Game{}

	ebiten.SetTPS(30)
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Game of life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
