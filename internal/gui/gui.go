package gui

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"mygame/internal/game"
)

type GameAdapter struct {
	game           *game.Game
	dudeRenderer   *DudeRenderer
	bulletRenderer *BulletRenderer
}

func (adpt *GameAdapter) Update() error {
	inputs := game.GameInput{
		LeftPressed:   ebiten.IsKeyPressed(ebiten.KeyA),
		RightPressed:  ebiten.IsKeyPressed(ebiten.KeyD),
		DownPressed:   ebiten.IsKeyPressed(ebiten.KeyS),
		UpPressed:     ebiten.IsKeyPressed(ebiten.KeyW),
		MouseLeftDown: ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft),
	}
	adpt.game.Update(&inputs)
	return nil
}

func (adpt *GameAdapter) Draw(screen *ebiten.Image) {
	for _, e := range adpt.game.Entities {
		switch t := e.(type) {
		case *game.Dude:
			adpt.dudeRenderer.Draw(t, screen)
		case *game.Bullet:
			adpt.bulletRenderer.Draw(t, screen)
		}
	}
}

func (adpt *GameAdapter) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 768
}

func RunGui(game *game.Game) {
	dudeRenderer := DudeRenderer{}
	if err := dudeRenderer.init(); err != nil {
		log.Fatalf("Failed to init. %v", err)
		os.Exit(1)
	}
	bulletRenderer := BulletRenderer{}
	if err := bulletRenderer.init(); err != nil {
		log.Fatalf("Failed to init. %v", err)
		os.Exit(1)
	}
	adapter := GameAdapter{
		game:           game,
		dudeRenderer:   &dudeRenderer,
		bulletRenderer: &bulletRenderer,
	}
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("mygame")
	if err := ebiten.RunGame(&adapter); err != nil {
		log.Fatal(err)
	}
}
