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
	audioPlayer    *AudioPlayer
}

func (adpt *GameAdapter) OnGameEvent(event game.GameEvent) {
	switch event.Type {
	case game.EventShoot:
		adpt.audioPlayer.PlayShootSound()
	}
}

type EventPrinter struct{}

func (ep *EventPrinter) OnGameEvent(event game.GameEvent) {
	println(event.Type)
}

func (adpt *GameAdapter) Update() error {
	cursorPosX, cursorPosY := ebiten.CursorPosition()
	mouseDown := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	inputs := game.GameInput{
		LeftPressed:   ebiten.IsKeyPressed(ebiten.KeyA),
		RightPressed:  ebiten.IsKeyPressed(ebiten.KeyD),
		DownPressed:   ebiten.IsKeyPressed(ebiten.KeyS),
		UpPressed:     ebiten.IsKeyPressed(ebiten.KeyW),
		MouseLeftDown: mouseDown,
		CursorPos:     game.Point2D{X: float64(cursorPosX), Y: float64(cursorPosY)},
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
	audioPlayer := AudioPlayer{}
	if err := audioPlayer.Init(); err != nil {
		log.Fatalf("Failed to init. %v", err)
		os.Exit(1)
	}
	adapter := GameAdapter{
		game:           game,
		dudeRenderer:   &dudeRenderer,
		bulletRenderer: &bulletRenderer,
		audioPlayer:    &audioPlayer,
	}

	game.Events.RegisterListener(&adapter)
	// game.Events.RegisterListener(&EventPrinter{})

	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("mygame")
	if err := ebiten.RunGame(&adapter); err != nil {
		log.Fatal(err)
	}
}
