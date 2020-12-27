package socket

import "realtime-calculator-api/socket/model"

type Hub interface {
	RegisteredClients() map[*model.Client]bool
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
