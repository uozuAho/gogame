package game

type Entity interface {
	Update(g *Game, input *GameInput)
}

type Game struct {
	DudeSpeedPerTick float32
	DudePos          Point2D
	Entities         []Entity
}

type GameInput struct {
	LeftPressed   bool
	RightPressed  bool
	UpPressed     bool
	DownPressed   bool
	MouseLeftDown bool
	CursorPos     Point2D
}

func NewGame() Game {
	entities := []Entity{
		&Dude{SpeedPerTick: 2, Pos: &Point2D{X: 300, Y: 300}, RespondToUserInput: true},
		&Dude{SpeedPerTick: 2, Pos: &Point2D{X: 500, Y: 500}},
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
