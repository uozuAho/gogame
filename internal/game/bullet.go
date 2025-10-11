package game

const BulletSpeedPerTick = 20

type Bullet struct {
	Pos      Point2D
	Velocity Point2D
}

func (d *Bullet) Update(g *Game, input *GameInput) {
	d.Pos.Add(d.Velocity)
}
