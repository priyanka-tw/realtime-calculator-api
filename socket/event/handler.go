package event

import (
	"errors"
	"realtime-calculator-api/socket"
	"realtime-calculator-api/socket/model"
)

type Count struct {
	numberOfUsers int
}

type Handler interface {
	Handle(currentClient *model.Client, data string) error
}

type EventHandler struct {
	hub socket.Hub
	*Count
}

func NewEventHandler(hub socket.Hub) EventHandler {
	return EventHandler{
		hub:   hub,
		Count: &Count{},
	}
}

func (eh EventHandler) getHandler(event string) (Handler, error) {
	switch event {
	case "login":
		return NewLoginHandler(eh.Count, eh.hub), nil
	default:
		return nil, errors.New("wrong event passed")
	}
}
