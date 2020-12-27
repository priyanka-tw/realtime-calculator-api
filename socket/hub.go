package socket

//go:generate mockgen -source=hub.go -destination=./mock/hub_mock.go -package=mock

import (
	"log"
	"realtime-calculator-api/socket/model"
)

type Hub interface {
	RegisteredClients() map[*model.Client]bool
	BroadcastToAllClients(handler model.EventMetadata) error
}

type hub struct {
	Clients map[*model.Client]bool
}

func NewHub() Hub {
	return &hub{Clients: make(map[*model.Client]bool)}
}
func (h *hub) RegisteredClients() map[*model.Client]bool {
	return h.Clients
}

func (h *hub) BroadcastToAllClients(metadata model.EventMetadata) error {
	log.Println("hub : Broadcast to all clients started")
	var broadcastError error
	for client := range h.Clients {
		err := client.Connection.WriteJSON(metadata)
		if err != nil {
			log.Println("error while broadcasting to a client, err: ", err)
			broadcastError = err
			continue
		}
	}
	log.Println("hub : Broadcast to all clients completed")
	return broadcastError
}
