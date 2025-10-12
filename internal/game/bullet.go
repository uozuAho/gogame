package game

const BulletSpeedPerTick = 20

type Bullet struct {
	Pos Point2D
	Dir Point2D // unit vec
}

func (d *Bullet) Update(g *Game, input *GameInput) {
	temp := d.Dir.Copy()
	temp.Multiply(BulletSpeedPerTick)
	d.Pos.Add(temp)
}

func NewBullet(px, py, dx, dy float64) *Bullet {
	return &Bullet{
		Pos: Point2D{X: px, Y: py},
		Dir: Point2D{X: dx, Y: dy}.UnitVec(),
	}
}
