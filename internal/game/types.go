package game

type Point2D struct {
	X int
	Y int
}

func (p1 *Point2D) Copy() Point2D {
	return Point2D{p1.X, p1.Y}
}

func (p1 *Point2D) Add(p2 Point2D) {
	p1.X += p2.X
	p1.Y += p2.Y
}

func (p1 *Point2D) Subtract(p2 Point2D) {
	p1.X -= p2.X
	p1.Y -= p2.Y
}

type Entity interface {
	Update(g *Game, input *GameInput)
}
