package socket

import "realtime-calculator-api/socket/model"

type CalculateHandler struct {
	hub Hub
}

func NewCalculateHandler(hub Hub) CalculateHandler {
	return CalculateHandler{hub: hub}
}

func (l CalculateHandler) Handle(currentClient *model.Client, data string) error {
	writeEventHandler := model.EventMetadata{
		Event: "history",
		Data:  data,
	}
	return l.hub.BroadcastToAllClients(writeEventHandler)
}
