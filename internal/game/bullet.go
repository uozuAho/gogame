package game

const BulletSpeedPerTick = 20

type Bullet struct {
	Pos Point2D
	Dir Point2D // unit vec
}

func (d *Bullet) Update(g *Game, input *GameInput) {
	// d.Pos.Add(d.Dir * 20)
}

// func NewBullet(start Point2D, dir Point2D) *Bullet {
// 	// todo: implement this
// 	return &Bullet{
// 		Pos: Point2D{X: start.X, Y: start.Y},
// 		Dir: Point2D{X: nx, Y: ny},
// 	}
// }
