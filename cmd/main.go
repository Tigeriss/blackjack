package main

import (
	"blackjack/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Render an img")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
