package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math"
)

type Game struct {
	gameMap *GameMap
	player *Player
}

func MakeGame() *Game {
	game := &Game{}
	game.gameMap = MakeMap("map.txt")
	game.player = &Player{x: 1, y: 3, hp: 100, atk: 1, def: 1}
	return game
}

func (g *Game) Update() error {
	g.player.input(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	winX, winY := screen.Size()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentTPS()), winX-80, winY-30)
	for y, row := range g.gameMap.gameMap {
		for x, tile := range row {
			if x == g.player.x && y == g.player.y {
				g.player.render(screen, x, y)
			} else {
				tile.render(screen, x, y, g)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	w := int(float64(outsideWidth) * s)
	h := int(float64(outsideHeight) * s)
	TileWidth = math.Min(float64(w), float64(h)) / 15
	return w, h
}
