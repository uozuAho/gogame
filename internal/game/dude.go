package game

type Dude struct {
	SpeedPerTick       float32
	Pos                *Point2D
	IsShooting         bool
	RespondToUserInput bool
}

func (d *Dude) Update(g *Game, input *GameInput) {
	if !d.RespondToUserInput {
		return
	}

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
	d.IsShooting = input.MouseLeftDown
}
