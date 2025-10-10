package game

type Point2D struct {
	X int
	Y int
}

type Entity interface {
	Update(g *Game, input *GameInput)
}

type Game struct {
	DudeSpeedPerTick float32
	DudePos          Point2D
	Entities         []Entity
}

type Dude struct {
	SpeedPerTick float32
	Pos          *Point2D
}

func (d *Dude) Update(g *Game, input *GameInput) {
	if input.DownPressed {
		d.Pos.Y += int(d.SpeedPerTick)
	}
	if input.UpPressed {
		d.Pos.Y -= int(d.SpeedPerTick)
	}
	if input.LeftPressed {
		d.Pos.X -= int(d.SpeedPerTick)
	}
	if input.RightPressed {
		d.Pos.X += int(d.SpeedPerTick)
	}
}

type GameInput struct {
	LeftPressed  bool
	RightPressed bool
	UpPressed    bool
	DownPressed  bool
}

func NewGame() Game {
	entities := []Entity{
		&Dude{SpeedPerTick: 2, Pos: &Point2D{X: 300, Y: 300}},
	}

	return Game{
		DudeSpeedPerTick: 2,
		DudePos:          Point2D{300, 300},
		Entities:         entities,
	}
}

func (g *Game) Update(input *GameInput) {
	for _, e := range g.Entities {
		e.Update(g, input)
	}
}
