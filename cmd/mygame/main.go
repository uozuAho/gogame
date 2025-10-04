package main

import (
	"mygame/internal/game"
	"mygame/internal/gui"
)

func main() {
	game := game.Game{}
	gui.RunGui(&game)
}
