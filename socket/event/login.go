package event

import (
	"realtime-calculator-api/socket"
	"realtime-calculator-api/socket/model"
	"strconv"
)

type LoginHandler struct {
	*Count
	hub socket.Hub
}

func NewLoginHandler(commonCount *Count, hub socket.Hub) LoginHandler {
	return LoginHandler{Count: commonCount, hub: hub}
}

func (l LoginHandler) Handle(currentClient *model.Client, data string) error {
	currentClient.Username = data
	l.numberOfUsers++

	writeEventHandler := model.EventMetadata{
		Event: "logged in users",
		Data:  strconv.Itoa(l.numberOfUsers),
	}
	return l.hub.BroadcastToAllClients(writeEventHandler)
}
