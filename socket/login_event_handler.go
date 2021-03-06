package socket

import (
	"log"
	"realtime-calculator-api/socket/model"
	"strconv"
)

type LoginHandler struct {
	*Count
	hub Hub
}

func NewLoginHandler(commonCount *Count, hub Hub) LoginHandler {
	return LoginHandler{Count: commonCount, hub: hub}
}

func (l LoginHandler) Handle(currentClient *model.Client, data string) error {
	log.Println("login event handler: handling event")
	currentClient.Username = data
	l.numberOfUsers++

	writeEventHandler := model.EventMetadata{
		Event: "logged in users",
		Data:  strconv.Itoa(l.numberOfUsers),
	}
	return l.hub.BroadcastToAllClients(writeEventHandler)
}
