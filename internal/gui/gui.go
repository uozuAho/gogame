package gui

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"mygame/internal/game"
)

type GameAdapter struct {
	game      *game.Game
	dudeImage *ebiten.Image
}

func (adpt *GameAdapter) Update() error {
	inputs := game.GameInput{
		LeftPressed:  ebiten.IsKeyPressed(ebiten.KeyA),
		RightPressed: ebiten.IsKeyPressed(ebiten.KeyD),
		DownPressed:  ebiten.IsKeyPressed(ebiten.KeyS),
		UpPressed:    ebiten.IsKeyPressed(ebiten.KeyW),
	}
	adpt.game.Update(&inputs)
	return nil
}

func (adpt *GameAdapter) Draw(screen *ebiten.Image) {
	if adpt.dudeImage == nil {
		ebitenutil.DebugPrint(screen, "missing dude image")
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(adpt.game.DudePos.X), float64(adpt.game.DudePos.Y))
	screen.DrawImage(adpt.dudeImage, op)
}

func (adpt *GameAdapter) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 768
}

func RunGui(game *game.Game) {
	img, _, err := ebitenutil.NewImageFromFile("assets/img/dude.png")
	if err != nil {
		log.Fatalf("failed to load dude image: %v", err)
	}

	adapter := GameAdapter{game: game, dudeImage: img}
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&adapter); err != nil {
		log.Fatal(err)
	}
}
