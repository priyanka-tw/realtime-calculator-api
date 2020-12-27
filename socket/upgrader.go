package socket

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type UpgraderWrapper interface {
	Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error)
}

type upgraderWrapper struct {
	upgrader *websocket.Upgrader
}

func NewUpgraderWrapper(upgrader *websocket.Upgrader) UpgraderWrapper {
	return upgraderWrapper{upgrader: upgrader}
}

func (a upgraderWrapper) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	return a.upgrader.Upgrade(w, r, responseHeader)
}
