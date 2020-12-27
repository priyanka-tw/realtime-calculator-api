package socket

import (
	"log"
	"realtime-calculator-api/socket/model"
)

type CalculateHandler struct {
	hub Hub
}

func NewCalculateHandler(hub Hub) CalculateHandler {
	return CalculateHandler{hub: hub}
}

func (l CalculateHandler) Handle(currentClient *model.Client, data string) error {
	log.Println("login event handler: handling event")

	writeEventHandler := model.EventMetadata{
		Event: "history",
		Data:  data,
	}
	return l.hub.BroadcastToAllClients(writeEventHandler)
}
