package game

const BulletSpeedPerTick = 20
const BulletDamage = 10

type Bullet struct {
	Dir     Point2D // unit vec
	Damage  float64
	Owner   *Dude
	topLeft Point2D
	isDead  bool
}

func (b *Bullet) Pos() Point2D {
	return b.topLeft
}

func (b *Bullet) Kill() {
	b.isDead = true
}

func (b *Bullet) IsDead() bool {
	return b.isDead
}

func (d *Bullet) Update(g *Game, input *GameInput) {
	temp := d.Dir.Copy()
	temp.Multiply(BulletSpeedPerTick)
	d.topLeft.Add(temp)
}

func NewBullet(px, py, dx, dy float64, owner *Dude) *Bullet {
	return &Bullet{
		Dir:     Point2D{X: dx, Y: dy}.UnitVec(),
		Damage:  BulletDamage,
		Owner:   owner,
		topLeft: Point2D{X: px, Y: py},
		isDead:  false,
	}
}
