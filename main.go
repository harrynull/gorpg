package main

import "log"
import "github.com/hajimehoshi/ebiten/v2"

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("GoRPG")

	if err := ebiten.RunGame(MakeGame()); err != nil {
		log.Fatal(err)
	}
}
