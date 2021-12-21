package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"strconv"
)

var TileWidth float64

type Tile interface {
	render(screen *ebiten.Image, x int, y int, game *Game)
	passable() bool
	action(game *Game) bool // return if should destruct
}

type EmptyTile struct {}
func (t *EmptyTile) render(_ *ebiten.Image, _ int, _ int, _ *Game) {}
func (t *EmptyTile) passable() bool { return true }
func (t *EmptyTile) action(game *Game) bool { return false }

type WallTile struct {}
func (t *WallTile) render(screen *ebiten.Image, x int, y int, _ *Game) {
	ebitenutil.DrawRect(
		screen,
		float64(x)*TileWidth,
		float64(y)*TileWidth,
		TileWidth,
		TileWidth,
		color.RGBA{A: 255},
	)
}
func (t *WallTile) passable() bool { return false }
func (t *WallTile) action(game *Game) bool { return false }

type ATKGemTile struct {}
func (t *ATKGemTile) render(screen *ebiten.Image, x int, y int, _ *Game) {
	ebitenutil.DrawRect(
		screen,
		float64(x)*TileWidth+TileWidth*.1,
		float64(y)*TileWidth+TileWidth*.1,
		TileWidth*.8,
		TileWidth*.8,
		color.RGBA{R:255, A: 255},
	)
}
func (t *ATKGemTile) passable() bool { return true }
func (t *ATKGemTile) action(game *Game) bool {
	game.player.atk++
	return true
}

type DEFGemTile struct {}
func (t *DEFGemTile) render(screen *ebiten.Image, x int, y int, _ *Game) {
	ebitenutil.DrawRect(
		screen,
		float64(x)*TileWidth+TileWidth*.1,
		float64(y)*TileWidth+TileWidth*.1,
		TileWidth*.8,
		TileWidth*.8,
		color.RGBA{B:255, A: 255},
	)
}
func (t *DEFGemTile) passable() bool { return true }
func (t *DEFGemTile) action(game *Game) bool {
	game.player.def++
	return true
}
type HealPotionTile struct {}
func (t *HealPotionTile) render(screen *ebiten.Image, x int, y int, _ *Game) {
	ebitenutil.DrawRect(
		screen,
		float64(x)*TileWidth+TileWidth*.1,
		float64(y)*TileWidth+TileWidth*.1,
		TileWidth*.8,
		TileWidth*.8,
		color.RGBA{G:255, A: 255},
	)
}
func (t *HealPotionTile) passable() bool { return true }
func (t *HealPotionTile) action(game *Game) bool {
	game.player.hp+=100
	return true
}

type MobTile struct {
	mobData Mob
}
func (t *MobTile) render(screen *ebiten.Image, x int, y int, game *Game) {
	ebitenutil.DrawRect(
		screen,
		float64(x)*TileWidth+TileWidth*.1,
		float64(y)*TileWidth+TileWidth*.1,
		TileWidth*.8,
		TileWidth*.8,
		color.RGBA{G: 150, A: 255},
	)
	text.Draw(
		screen,
		strconv.Itoa(t.mobData.getDmg(game.player)),
		dmgFont,
		int(float64(x)*TileWidth+TileWidth*.1),
		int(float64(y)*TileWidth+TileWidth*.8),
		color.Black)
}
func (t *MobTile) passable() bool { return true }
func (t *MobTile) action(game *Game) bool {
	game.player.hp -= t.mobData.getDmg(game.player)
	return true
}

func MakeTileFromChar(c string, mobsMap *map[string]Mob) Tile {
	switch c {
	case " ":
		return &EmptyTile{}
	case "x":
		return &WallTile{}
	case "a":
		return &ATKGemTile{}
	case "d":
		return &DEFGemTile{}
	case "h":
		return &HealPotionTile{}
	}
	if mob, ok := (*mobsMap)[c]; ok {
		return &MobTile{mob}
	} else {
		println(c, "not recognized")
		return nil
	}
}
