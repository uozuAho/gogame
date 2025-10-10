package gui

import (
	"log"
	"mygame/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type DudeRenderer struct {
	dudeImage *ebiten.Image
}

func (rdr *DudeRenderer) init() error {
	img, _, err := ebitenutil.NewImageFromFile("assets/img/dude.png")
	if err != nil {
		log.Fatalf("failed to load dude image: %v", err)
		return err
	}
	rdr.dudeImage = img
	return nil
}

func (rdr *DudeRenderer) DrawDude(dude *game.Dude, screen *ebiten.Image) {
	if rdr.dudeImage == nil {
		ebitenutil.DebugPrint(screen, "missing dude image")
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(dude.Pos.X), float64(dude.Pos.Y))
	screen.DrawImage(rdr.dudeImage, op)
}
