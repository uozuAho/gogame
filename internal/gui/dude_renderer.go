package gui

import (
	"image/color"
	"log"
	"mygame/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

	if dude.IsShooting {
		mx, my := ebiten.CursorPosition()

		// compute dude center
		w := rdr.dudeImage.Bounds().Dx()
		h := rdr.dudeImage.Bounds().Dy()
		dx := float64(dude.Pos.X) + float64(w)/2
		dy := float64(dude.Pos.Y) + float64(h)/2

		// draw red line using vector.StrokeLine (no anti-aliasing)
		vector.StrokeLine(screen, float32(dx), float32(dy), float32(mx), float32(my), 2.0, color.RGBA{R: 255, G: 0, B: 0, A: 255}, false)
	}
}
