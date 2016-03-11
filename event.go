package gonsole

import "fmt"

type EventSource interface {
	ID() string
}

type Event struct {
	Type   string
	Source EventSource
	Data   map[string]interface{}
}

type EventDispatcher struct {
	registeredEvents map[string][]func(ev *Event) bool
}

func NewEventDispatcher() *EventDispatcher {
	ed := &EventDispatcher{
		registeredEvents: make(map[string][]func(ev *Event) bool, 0),
	}
	return ed
}

func (ed *EventDispatcher) SubmitEvent(ev *Event) {
	key := ed.getKey(ev.Source, ev.Type)
	if funcs, ok := ed.registeredEvents[key]; ok {
		for _, function := range funcs {
			function(ev)
		}
	}
}

func (ed *EventDispatcher) AddEventListener(source EventSource, eventType string, handler func(ev *Event) bool) {
	key := ed.getKey(source, eventType)
	funcArray, ok := ed.registeredEvents[key]
	if !ok {
		funcArray = make([]func(ev *Event) bool, 0)
	}
	funcArray = append(funcArray, handler)
	ed.registeredEvents[key] = funcArray
}

func (ed *EventDispatcher) getKey(source EventSource, eventType string) string {
	return fmt.Sprintf("%s___%s", source.ID(), eventType)
}
