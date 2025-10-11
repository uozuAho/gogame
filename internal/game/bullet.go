package game

import "math"

const BulletSpeedPerTick = 20

type Bullet struct {
	Pos      Point2D
	Velocity Point2D
}

func (d *Bullet) Update(g *Game, input *GameInput) {
	d.Pos.Add(d.Velocity)
}

// NewBullet creates a bullet at `start` travelling in the direction of `dir`.
// `dir` may be a simple integer direction (e.g. {1,0} for right); it will be
// normalized and scaled by BulletSpeedPerTick.
func NewBullet(start Point2D, dir Point2D) *Bullet {
	dx := float64(dir.X)
	dy := float64(dir.Y)
	mag := math.Hypot(dx, dy)
	if mag == 0 {
		// Default to the right if no direction provided
		dx = 1
		mag = 1
	}
	nx := dx / mag
	ny := dy / mag
	vx := int(math.Round(nx * float64(BulletSpeedPerTick)))
	vy := int(math.Round(ny * float64(BulletSpeedPerTick)))

	return &Bullet{
		Pos:      Point2D{X: start.X, Y: start.Y},
		Velocity: Point2D{X: vx, Y: vy},
	}
}
