package gui

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"mygame/internal/game"
)

type GameAdapter struct {
	game *game.Game
}

func (adpt *GameAdapter) Update() error {
	inputs := game.GameInput{
		LeftPressed:  ebiten.IsKeyPressed(ebiten.KeyLeft),
		RightPressed: ebiten.IsKeyPressed(ebiten.KeyRight),
		DownPressed:  ebiten.IsKeyPressed(ebiten.KeyDown),
		UpPressed:    ebiten.IsKeyPressed(ebiten.KeyUp),
	}
	adpt.game.Update(&inputs)
	return nil
}

func (adpt *GameAdapter) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (adpt *GameAdapter) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 768
}

func RunGui(game *game.Game) {
	adapter := GameAdapter{game}
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&adapter); err != nil {
		log.Fatal(err)
	}
}
