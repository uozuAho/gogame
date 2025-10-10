package gui

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"mygame/internal/game"
)

type GameAdapter struct {
	game         *game.Game
	dudeRenderer *DudeRenderer
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
	for _, e := range adpt.game.Entities {
		switch t := e.(type) {
		case *game.Dude:
			adpt.dudeRenderer.DrawDude(t, screen)
		}
	}

	// todo: reimplement this
	// // Draw a line from the center of the dude to the mouse cursor
	// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
	// 	// get cursor position
	// 	mx, my := ebiten.CursorPosition()

	// 	// compute dude center
	// 	w := adpt.dudeImage.Bounds().Dx()
	// 	h := adpt.dudeImage.Bounds().Dy()
	// 	dx := float64(adpt.game.DudePos.X) + float64(w)/2
	// 	dy := float64(adpt.game.DudePos.Y) + float64(h)/2

	// 	// draw red line using vector.StrokeLine (no anti-aliasing)
	// 	vector.StrokeLine(screen, float32(dx), float32(dy), float32(mx), float32(my), 2.0, color.RGBA{R: 255, G: 0, B: 0, A: 255}, false)
	// }
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
	adapter := GameAdapter{game: game, dudeRenderer: &dudeRenderer}
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&adapter); err != nil {
		log.Fatal(err)
	}
}
