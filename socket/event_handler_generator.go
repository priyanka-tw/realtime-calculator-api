package socket

//go:generate mockgen -source=event_handler_generator.go -destination=./mock/event_handler_generator_mock.go =package=mock

import (
	"errors"
	"log"
	"realtime-calculator-api/socket/interface"
)

type Count struct {
	numberOfUsers int
}

type EventHandlerGenerator interface {
	GetHandler(event string) (_interface.EventHandler, error)
}

type eventHandlerGenerator struct {
	hub Hub
	*Count
}

func NewEventHandlerGenerator(hub Hub) EventHandlerGenerator {
	return eventHandlerGenerator{
		hub:   hub,
		Count: &Count{},
	}
}

func (eh eventHandlerGenerator) GetHandler(event string) (_interface.EventHandler, error) {
	log.Println("event handler generator: get handler")
	switch event {
	case "login":
		return NewLoginHandler(eh.Count, eh.hub), nil
	case "logout":
		return NewLogoutHandler(eh.Count, eh.hub), nil
	case "calculate":
		return NewCalculateHandler(eh.hub), nil
	default:
		return nil, errors.New("wrong event passed")
	}
}
