### Phase 1: Create Event System Infrastructure

**File: internal/game/events.go** (new file)

```go package game

// EventType identifies different game events type EventType int

const (
    EventShoot EventType = iota // Future: EventHit, EventDeath, EventSpawn,
    etc.
)

// GameEvent represents something that happened in the game type GameEvent
struct {
    Type     EventType EntityID string // or uintptr, or Entity reference Data
    interface{} // event-specific data
}

// EventListener is the interface for components that want to receive events
type EventListener interface {
    OnGameEvent(event GameEvent)
}

// EventEmitter collects events during a game tick type EventEmitter struct {
    events    []GameEvent listeners []EventListener
}

func (e *EventEmitter) EmitEvent(event GameEvent) {
    e.events = append(e.events, event)
}

func (e *EventEmitter) RegisterListener(listener EventListener) {
    e.listeners = append(e.listeners, listener)
}

func (e *EventEmitter) DispatchEvents() {
    for _, event := range e.events {
        for _, listener := range e.listeners {
            listener.OnGameEvent(event)
        }
    } // Clear events after dispatch e.events = e.events[:0]
}
```

### Phase 2: Integrate Event System into Game

**Modify: internal/game/game.go**

```go
type Game struct {
    Entities []Entity
    Events   *EventEmitter  // Add event system
}

func NewGame() Game {
    entities := []Entity{
        &Dude{SpeedPerTick: 2, Pos: Point2D{X: 300, Y: 300}, RespondToUserInput:
        true}, &Dude{SpeedPerTick: 2, Pos: Point2D{X: 500, Y: 500}},
    }

    return Game{
        Entities: entities,
        Events:   &EventEmitter{},  // Initialize event      system
    }
}

func (g *Game) Update(input *GameInput) {
    // Update all entities (they may emit events)
    for _, e := range g.Entities {
        e.Update(g, input)
    }

    // Dispatch events after all updates complete
    // // This ensures consistent  ordering and allows GUI to react g.Events.DispatchEvents()
}
```

### Phase 3: Entities Emit Events

**Modify: internal/game/dude.go**

In `Dude.Update()`, replace the implicit shooting action with explicit event
emission:

```go func (dude *Dude) Update(g *Game, input *GameInput) {
    // ... existing movement code ...

    if input.MouseLeftDown && !dude.prevMouseLeftDown {
        dir := input.CursorPos.Copy() dir.Subtract(dude.Pos) b :=
        NewBullet(dude.Pos.X, dude.Pos.Y, dir.X, dir.Y) g.Entities =
        append(g.Entities, b)

        // EMIT SHOOT EVENT - this is the key change
        g.Events.EmitEvent(GameEvent{
            Type:     EventShoot,
            EntityID: fmt.Sprintf("%p", dude), // or use a proper ID system
            Data:     nil, // could include position,     direction, etc.
        })
    }

    dude.IsShooting = input.MouseLeftDown dude.prevMouseLeftDown =
    input.MouseLeftDown
}
```

### Phase 4: GUI Responds to Events

**Modify: internal/gui/gui.go**

Make GameAdapter implement EventListener and respond to shoot events:

```go
// Implement EventListener interface func (adpt *GameAdapter)
OnGameEvent(event game.GameEvent) {
    switch event.Type { case game.EventShoot:
        adpt.playShootSound()
    // Future: handle other event types }
}

// Extract sound playing to separate method func (adpt *GameAdapter)
playShootSound() {
    if adpt.audioCtx == nil || len(adpt.shootWav) == 0 {
        return
    }

    // Play the shot sound in a goroutine so we don't block go func() {
        r := bytes.NewReader(adpt.shootWav) s, err := wav.Decode(adpt.audioCtx,
        r) if err != nil {
            log.Printf("failed to decode shoot wav: %v", err) return
        } player, err := audio.NewPlayer(adpt.audioCtx, s) if err != nil {
            log.Printf("failed to create audio player: %v", err) return
        } player.Play()
    }()
}

func (adpt *GameAdapter) Update() error {
    // Get input cursorPosX, cursorPosY := ebiten.CursorPosition() mouseDown :=
    ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

    // REMOVE the mouse edge detection and sound playing code here

    inputs := game.GameInput{
        LeftPressed:   ebiten.IsKeyPressed(ebiten.KeyA), RightPressed:
        ebiten.IsKeyPressed(ebiten.KeyD), DownPressed:
        ebiten.IsKeyPressed(ebiten.KeyS), UpPressed:
        ebiten.IsKeyPressed(ebiten.KeyW), MouseLeftDown: mouseDown, CursorPos:
        game.Point2D{X: float64(cursorPosX), Y: float64(cursorPosY)},
    }

    // Update game (events will be dispatched internally)
    adpt.game.Update(&inputs)

    return nil
}

// Register as event listener during initialization func RunGui(game *game.Game)
{
    // ... existing initialization ...

    adapter := GameAdapter{
        game:           game,
        dudeRenderer:   &dudeRenderer,
        bulletRenderer: &bulletRenderer,
    }

    // ... audio setup ...

    // Register adapter as event listener game.Events.RegisterListener(&adapter)

    // ... rest of initialization ...
}
```

### Phase 5: Remove Obsolete Code

Remove from GameAdapter: - `prevMouseLeftDown` field (no longer needed in GUI) -
Mouse edge detection logic in Update() (lines 28-44 of current gui.go)
