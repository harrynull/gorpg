package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

var (
	normalFont font.Face
	dmgFont font.Face
)

type Player struct {
	x, y int
	hp int
	atk int
	def int
}

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 128
	normalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    20,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	dmgFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    10,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 20
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func (p *Player) input(game *Game) {
	var moved = false
	if repeatingKeyPressed(ebiten.KeyArrowUp) && !game.gameMap.collisionCheck(p.x, p.y-1) {
		p.y--
		moved = true
	}
	if repeatingKeyPressed(ebiten.KeyArrowLeft) && !game.gameMap.collisionCheck(p.x-1, p.y) {
		p.x--
		moved = true
	}
	if repeatingKeyPressed(ebiten.KeyArrowRight) && !game.gameMap.collisionCheck(p.x+1, p.y) {
		p.x++
		moved = true
	}
	if repeatingKeyPressed(ebiten.KeyArrowDown) && !game.gameMap.collisionCheck(p.x, p.y+1) {
		p.y++
		moved = true
	}
	if moved {
		if game.gameMap.gameMap[p.y][p.x].action(game) {
			game.gameMap.gameMap[p.y][p.x] = &EmptyTile{}
		}
	}
}

func (p *Player) render(screen *ebiten.Image, x int, y int) {
	ebitenutil.DrawRect(
		screen,
		float64(x)*TileWidth+TileWidth*.1,
		float64(y)*TileWidth+TileWidth*.1,
		TileWidth*0.8,
		TileWidth*0.8,
		color.RGBA{G: 100, A: 255},
	)
	text.Draw(screen, fmt.Sprintf("HP: %d", p.hp), normalFont, int(TileWidth*11), int(TileWidth), color.Black)
	text.Draw(screen, fmt.Sprintf("ATK: %d", p.atk), normalFont, int(TileWidth*11), int(TileWidth*2), color.Black)
	text.Draw(screen, fmt.Sprintf("DEF: %d", p.def), normalFont, int(TileWidth*11), int(TileWidth*3), color.Black)
}