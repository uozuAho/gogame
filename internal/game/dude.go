package game

import "fmt"

type Dude struct {
	SpeedPerTick       float64
	IsShooting         bool
	RespondToUserInput bool
	prevMouseLeftDown  bool
	HitPoints          float64
	MaxHitPoints       float64
	topLeft            Point2D
}

func NewDude(pos Point2D) *Dude {
	return &Dude{
		SpeedPerTick: 2,
		topLeft:      pos,
		HitPoints:    100,
		MaxHitPoints: 100,
	}
}

func (dude *Dude) Pos() Point2D {
	return dude.topLeft
}

func (dude *Dude) DoDamage(damage float64) {
	dude.HitPoints -= damage
}

func (dude *Dude) Update(g *Game, input *GameInput) {
	if !dude.RespondToUserInput {
		return
	}

	if input.DownPressed {
		dude.topLeft.Y += dude.SpeedPerTick
	}
	if input.UpPressed {
		dude.topLeft.Y -= dude.SpeedPerTick
	}
	if input.LeftPressed {
		dude.topLeft.X -= dude.SpeedPerTick
	}
	if input.RightPressed {
		dude.topLeft.X += dude.SpeedPerTick
	}

	if input.MouseLeftDown && !dude.prevMouseLeftDown {
		dir := input.CursorPos.Copy()
		dir.Subtract(dude.topLeft)
		b := NewBullet(dude.topLeft.X, dude.topLeft.Y, dir.X, dir.Y)
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
