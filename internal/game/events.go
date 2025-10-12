package game

type EventType int

const (
	EventShoot EventType = iota
)

type GameEvent struct {
	Type     EventType
	EntityID string
	Data     interface{}
}

type EventListener interface {
	OnGameEvent(event GameEvent)
}

type EventEmitter struct {
	events    []GameEvent
	listeners []EventListener
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
	}
	e.events = e.events[:0]
}
