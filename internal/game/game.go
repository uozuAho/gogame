package game

type Entity interface {
	Update(g *Game, input *GameInput)
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
	tempPlayer := &Dude{SpeedPerTick: 2, Pos: Point2D{X: 300, Y: 300}, RespondToUserInput: true}

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
	g.Events.DispatchEvents()
}
