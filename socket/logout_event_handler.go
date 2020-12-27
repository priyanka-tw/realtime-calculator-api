package socket

import (
	"log"
	"realtime-calculator-api/socket/model"
	"strconv"
)

type LogoutHandler struct {
	*Count
	hub Hub
}

func NewLogoutHandler(commonCount *Count, hub Hub) LogoutHandler {
	return LogoutHandler{
		Count: commonCount,
		hub:   hub,
	}
}

func (l LogoutHandler) Handle(currentClient *model.Client, data string) error {
	log.Println("logout event handler: handling event")

	l.numberOfUsers--
	writeEventHandler := model.EventMetadata{
		Event: "logged in users",
		Data:  strconv.Itoa(l.numberOfUsers),
	}
	registeredClients := l.hub.RegisteredClients()
	delete(registeredClients, currentClient)

	return l.hub.BroadcastToAllClients(writeEventHandler)
}
