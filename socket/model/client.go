package model

type Conn interface {
	WriteJSON(v interface{}) error
	ReadJSON(v interface{}) error
	Close() error
}

type Client struct {
	Connection Conn
	Username string
}
