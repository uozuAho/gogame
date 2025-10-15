package game

const BulletSpeedPerTick = 20
const BulletDamage = 10

type Bullet struct {
	Dir     Point2D // unit vec
	Damage  float64
	topLeft Point2D
}

func (b *Bullet) Pos() Point2D {
	return b.topLeft
}

func (d *Bullet) Update(g *Game, input *GameInput) {
	temp := d.Dir.Copy()
	temp.Multiply(BulletSpeedPerTick)
	d.topLeft.Add(temp)
}

func NewBullet(px, py, dx, dy float64) *Bullet {
	return &Bullet{
		topLeft: Point2D{X: px, Y: py},
		Dir:     Point2D{X: dx, Y: dy}.UnitVec(),
		Damage:  BulletDamage,
	}
}
