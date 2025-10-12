package gui

import (
	"bytes"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"

	"mygame/internal/game"
)

type GameAdapter struct {
	game              *game.Game
	dudeRenderer      *DudeRenderer
	bulletRenderer    *BulletRenderer
	audioCtx          *audio.Context
	shootWav          []byte
	prevMouseLeftDown bool
}

func (adpt *GameAdapter) Update() error {
	// detect mouse click edge so we can play a sound when the user shoots
	cursorPosX, cursorPosY := ebiten.CursorPosition()
	mouseDown := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if mouseDown && !adpt.prevMouseLeftDown && adpt.audioCtx != nil && len(adpt.shootWav) > 0 {
		// Play the shot sound in a goroutine so we don't block the UI/update loop
		go func() {
			r := bytes.NewReader(adpt.shootWav)
			s, err := wav.Decode(adpt.audioCtx, r)
			if err != nil {
				log.Printf("failed to decode shoot wav: %v", err)
				return
			}
			player, err := audio.NewPlayer(adpt.audioCtx, s)
			if err != nil {
				log.Printf("failed to create audio player: %v", err)
				return
			}
			player.Play()
		}()
	}

	inputs := game.GameInput{
		LeftPressed:   ebiten.IsKeyPressed(ebiten.KeyA),
		RightPressed:  ebiten.IsKeyPressed(ebiten.KeyD),
		DownPressed:   ebiten.IsKeyPressed(ebiten.KeyS),
		UpPressed:     ebiten.IsKeyPressed(ebiten.KeyW),
		MouseLeftDown: mouseDown,
		CursorPos:     game.Point2D{X: float64(cursorPosX), Y: float64(cursorPosY)},
	}
	adpt.game.Update(&inputs)
	adpt.prevMouseLeftDown = mouseDown
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
	const sampleRate = 44100
	adapter.audioCtx = audio.NewContext(sampleRate)
	data, err := os.ReadFile("assets/audio/white-short.wav")
	if err != nil {
		log.Fatalf("Failed to load shoot wav: %v", err)
		os.Exit(1)
	}
	adapter.shootWav = data
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("mygame")
	if err := ebiten.RunGame(&adapter); err != nil {
		log.Fatal(err)
	}
}
