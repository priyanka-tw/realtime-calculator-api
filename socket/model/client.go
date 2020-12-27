package model

type Conn interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteJSON(v interface{}) error
	Close() error
}

type Client struct {
	Connection Conn
	Username string
}
