package game

type Dude struct {
	SpeedPerTick       float32
	Pos                *Point2D
	IsShooting         bool
	RespondToUserInput bool
	// previous input state to detect rising-edge of shooting (first frame pressed)
	prevMouseLeftDown bool
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
	// rising-edge detection: if mouse left is now down but wasn't last update,
	// spawn a bullet in the current facing direction (use right if ambiguous).
	if input.MouseLeftDown && !d.prevMouseLeftDown {
		// Determine direction. For simplicity use the movement keys to infer
		// direction; if none pressed, default to right (1,0).
		dir := Point2D{X: 1, Y: 0}
		if input.LeftPressed {
			dir = Point2D{X: -1, Y: 0}
		} else if input.RightPressed {
			dir = Point2D{X: 1, Y: 0}
		} else if input.UpPressed {
			dir = Point2D{X: 0, Y: -1}
		} else if input.DownPressed {
			dir = Point2D{X: 0, Y: 1}
		}

		// Spawn bullet at dude's position
		if d.Pos != nil {
			b := NewBullet(Point2D{X: d.Pos.X, Y: d.Pos.Y}, dir)
			g.Entities = append(g.Entities, b)
		}
	}

	d.IsShooting = input.MouseLeftDown
	d.prevMouseLeftDown = input.MouseLeftDown
}
