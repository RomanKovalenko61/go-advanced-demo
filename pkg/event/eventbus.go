package event

const (
	EventLinkVisited = "link.visited"
)

type Event struct {
	Type string
	Data any
}

type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus: make(chan Event),
	}
}

func (e *EventBus) Publish(evevt Event) {
	e.bus <- evevt
}

func (e *EventBus) Subscribe() <-chan Event {
	return e.bus
}
