package game

type Entity interface {
	Update(g *Game, input *GameInput)
	Pos() Point2D
}

type Game struct {
	Entities []Entity
	Events   *EventEmitter
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
	tempPlayer := NewDude(Point2D{X: 300, Y: 300})
	tempPlayer.RespondToUserInput = true

	entities := []Entity{
		tempPlayer,
		NewDude(Point2D{X: 500, Y: 500}),
		NewDude(Point2D{X: 10, Y: 10}),
	}

	return Game{
		Entities: entities,
		Events:   &EventEmitter{},
	}
}

func (g *Game) Update(input *GameInput) {
	for _, e := range g.Entities {
		e.Update(g, input)
	}
	g.CheckCollisions()
	g.Events.DispatchEvents()
}

func (g *Game) CheckCollisions() {
	for i := 0; i < len(g.Entities)-1; i++ {
		for j := i + 1; j < len(g.Entities); j++ {
			ei := g.Entities[i]
			ej := g.Entities[j]

			// todo: extract const
			if ei.Pos().DistanceTo(ej.Pos()) < 20 {
				eib, eiIsBullet := ei.(*Bullet)
				ejb, ejIsBullet := ej.(*Bullet)
				eip, eiIsPlayer := ei.(*Dude)
				ejp, ejIsPlayer := ej.(*Dude)

				if eiIsBullet && ejIsPlayer {
					g.Events.EmitEvent(GameEvent{Type: EventCollision, EntityID: "", Data: ""})
					ejp.DoDamage(eib.Damage)
				} else if eiIsPlayer && ejIsBullet {
					g.Events.EmitEvent(GameEvent{Type: EventCollision, EntityID: "", Data: ""})
					eip.DoDamage(ejb.Damage)
				}
			}
		}
	}
}
