# Analysis of Current Architecture Issue

## Problem Statement

Currently, gui.go (internal/gui/gui.go:28-44) detects mouse click edges to play
the shoot sound. This violates several software design principles:

### Current Flow: 1. GUI detects mouse click edge (gui.go:28) 2. GUI plays sound
immediately (gui.go:30-43) 3. GUI passes input to Game.Update() (gui.go:54) 4.
Dude detects the same mouse click edge (dude.go:29) 5. Dude creates a bullet
(dude.go:32-33)

### Issues with Current Design:

1. **Violation of Single Responsibility Principle**: The GUI layer is
responsible for both rendering/input handling AND game event audio feedback.

2. **Tight Coupling**: The GUI has intimate knowledge of game semantics (what a
mouse click means in game terms - shooting).

3. **Duplicate Logic**: Both GUI (gui.go:28) and Dude (dude.go:29) detect the
mouse click edge independently using the same pattern.

4. **Lack of Separation of Concerns**: The presentation layer (GUI) is making
decisions about game events rather than just presenting what the game tells it
to present.

5. **Poor Extensibility**:
   - What if multiple dudes can shoot? Do we play the sound multiple times? -
   What if shooting requires ammo checks or cooldowns? - What if different
   entities make different sounds? - The GUI would need to know all this game
   logic.

6. **Breaks Determinism Goal**: The GUI layer playing sounds based on raw input
rather than game state means the audio is not truly coupled to what actually
happened in the game. If the game state prevents shooting (e.g., out of ammo),
the sound would still play.

7. **Headless Mode Coupling**: The GUI shouldn't need to understand game events
- it makes the abstraction between headless and GUI modes unclear.

---

## Recommended Architecture: Event-Driven Audio System

### Core Pattern: Observer/Event System

Implement a game event system where: - Game entities emit events when
significant actions occur - GUI subscribes to relevant events and responds with
appropriate presentation (audio, visual effects, etc.) - Game logic remains
completely decoupled from presentation

### Design Principles Applied:

1. **Single Responsibility**: Game handles logic, GUI handles presentation 2.
**Dependency Inversion**: Game depends on abstractions (event emitters), not
concrete GUI code 3. **Open/Closed**: Easy to add new event types without
modifying existing code 4. **Tell, Don't Ask**: Entities tell the game what
happened; GUI doesn't ask/detect

---

## Proposed Implementation Plan

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

---

## Benefits of This Design

1. **True Separation of Concerns**: Game logic is pure and knows nothing about
presentation 2. **Deterministic**: Audio is triggered by game events, not input
detection in GUI 3. **Headless Compatible**: Event system exists whether GUI is
present or not 4. **Extensible**: Easy to add new events (hits, deaths, spawns)
and new listeners (particle effects, screen shake, etc.) 5. **Testable**: Game
logic can be tested without GUI; audio can be tested by mocking EventListener 6.
**Flexible**: Multiple listeners can respond to the same event (e.g., sound +
visual effect) 7. **Maintainable**: Clear ownership of responsibilities

---

## Alternative Patterns Considered

### Alternative 1: Callback Functions Pass audio callbacks to Game/Dude during
construction. **Rejected**: Creates awkward dependencies and doesn't scale to
multiple listeners.

### Alternative 2: Return Events from Update() Have Update() return a list of
events. **Rejected**: Requires changing the Entity interface and makes event
collection more complex.

### Alternative 3: Polling Game State GUI checks game state each frame for
changes (e.g., check bullet count). **Rejected**: Inefficient, error-prone,
requires GUI to understand game state structure.

---

## Future Enhancements

Once the event system is in place:

1. **Rich Event Data**: Include position, entity references, damage amounts,
etc. 2. **Event Filtering**: Listeners can subscribe to specific event types 3.
**Event History**: Record events for replay/debugging 4. **Visual Effects**: Add
particle system that listens to events 5. **Sound Variation**: Different
entities or weapons make different sounds 6. **Network Events**: Events become
the basis for network synchronization 7. **Scripting**: User scripts can listen
to events and react (aligns with project goal)

---

## Implementation Risks and Mitigations

### Risk 1: Event System Overhead **Mitigation**: Pre-allocate event slices,
clear rather than reallocate, profile if needed

### Risk 2: Event Ordering Issues **Mitigation**: Document that events are
dispatched after all entity updates, maintain deterministic iteration order

### Risk 3: Breaking Headless Mode **Mitigation**: EventEmitter works without
listeners; test headless mode explicitly

### Risk 4: Memory Management with Goroutines **Mitigation**: Current sound
goroutine pattern is already in use; no change here

---

## Summary

The current architecture violates separation of concerns by having the GUI layer
detect and respond to game semantics (shooting). The proposed event-driven
system cleanly separates game logic from presentation, aligns with the project's
determinism and headless goals, and provides a scalable foundation for future
features. The implementation can be done incrementally without breaking existing
functionality.







Key Findings

  The current architecture has the GUI detecting mouse clicks to play the shoot sound (internal/gui/gui.go:28-44), which
  violates separation of concerns. The GUI shouldn't understand game semantics like "shooting" - it should only present what
  the game tells it to present.

  Proposed Solution

  Implement an event-driven architecture using the Observer pattern:

  1. Event System: Game entities emit events when significant actions occur (e.g., EventShoot)
  2. Game Controls Events: The Game.Update() collects and dispatches events after entity updates
  3. GUI Listens: GameAdapter implements EventListener interface and responds to events with appropriate audio/visual feedback
  4. Clean Separation: Game logic remains pure and presentation-agnostic

  This design:
  - Maintains your determinism goal (audio tied to actual game events, not raw input)
  - Supports headless mode (events work whether GUI is listening or not)
  - Enables future scripting (scripts can subscribe to events)
  - Follows SOLID principles (Single Responsibility, Dependency Inversion, Open/Closed)
  - Makes the codebase more maintainable and extensible

  The plan includes detailed implementation steps, code examples for each phase, analysis of alternatives, and mitigation
  strategies for potential risks.
