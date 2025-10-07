package main

import (
	"mygame/internal/game"
	"mygame/internal/gui"
)

func main() {
	game := game.NewGame()
	gui.RunGui(&game)
}
