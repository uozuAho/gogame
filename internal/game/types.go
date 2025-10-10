package game

type Point2D struct {
	X int
	Y int
}

type Entity interface {
	Update(g *Game, input *GameInput)
}
