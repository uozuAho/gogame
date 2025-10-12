package game

import "fmt"

type Dude struct {
	SpeedPerTick       float64
	Pos                Point2D
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
		dir.Subtract(dude.Pos)
		b := NewBullet(dude.Pos.X, dude.Pos.Y, dir.X, dir.Y)
		g.Entities = append(g.Entities, b)

		g.Events.EmitEvent(GameEvent{
			Type:     EventShoot,
			EntityID: fmt.Sprintf("%p", dude),
			Data:     nil,
		})
	}

	dude.IsShooting = input.MouseLeftDown
	dude.prevMouseLeftDown = input.MouseLeftDown
}
