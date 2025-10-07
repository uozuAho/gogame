package game

type Point2D struct {
	X int
	Y int
}

type Game struct {
	DudeSpeedPerTick float32
	DudePos          Point2D
}

type GameInput struct {
	LeftPressed  bool
	RightPressed bool
	UpPressed    bool
	DownPressed  bool
}

func NewGame() Game {
	return Game{
		DudeSpeedPerTick: 1,
		DudePos:          Point2D{300, 300},
	}
}

func (g *Game) Update(input *GameInput) {
	if input.DownPressed {
		g.DudePos.Y += int(g.DudeSpeedPerTick)
	}
	if input.UpPressed {
		g.DudePos.Y -= int(g.DudeSpeedPerTick)
	}
	if input.LeftPressed {
		g.DudePos.X += int(g.DudeSpeedPerTick)
	}
	if input.RightPressed {
		g.DudePos.X -= int(g.DudeSpeedPerTick)
	}
}
