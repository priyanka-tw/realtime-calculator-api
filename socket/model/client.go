package model

type Conn interface {
	ReadMessage() (messageType int, p []byte, err error)
	Close() error
}

type Client struct {
	Connection Conn
}
