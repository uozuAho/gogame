package game

type Dude struct {
	SpeedPerTick       float64
	Pos                *Point2D
	IsShooting         bool
	RespondToUserInput bool
	prevMouseLeftDown  bool
}

func (dude *Dude) Update(g *Game, input *GameInput) {
	if !dude.RespondToUserInput {
		return
	}

	if input.DownPressed {
		dude.Pos.Y += dude.SpeedPerTick
	}
	if input.UpPressed {
		dude.Pos.Y -= dude.SpeedPerTick
	}
	if input.LeftPressed {
		dude.Pos.X -= dude.SpeedPerTick
	}
	if input.RightPressed {
		dude.Pos.X += dude.SpeedPerTick
	}

	if input.MouseLeftDown && !dude.prevMouseLeftDown {
		dir := input.CursorPos.Copy()
		dir.Subtract(g.DudePos)

		// if dude.Pos != nil {
		// 	b := NewBullet(Point2D{X: dude.Pos.X, Y: dude.Pos.Y}, dir)
		// 	g.Entities = append(g.Entities, b)
		// }
	}

	dude.IsShooting = input.MouseLeftDown
	dude.prevMouseLeftDown = input.MouseLeftDown
}
