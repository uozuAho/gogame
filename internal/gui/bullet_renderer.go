package gui

import (
	"log"
	"mygame/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type BulletRenderer struct {
	bulletImage *ebiten.Image
}

func (rdr *BulletRenderer) init() error {
	img, _, err := ebitenutil.NewImageFromFile("assets/img/bullet.png")
	if err != nil {
		log.Fatalf("failed to load bullet image: %v", err)
		return err
	}
	rdr.bulletImage = img
	return nil
}

func (rdr *BulletRenderer) Draw(bullet *game.Bullet, screen *ebiten.Image) {
	if rdr.bulletImage == nil {
		ebitenutil.DebugPrint(screen, "missing dude image")
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(bullet.Pos.X), float64(bullet.Pos.Y))
	screen.DrawImage(rdr.bulletImage, op)
}
