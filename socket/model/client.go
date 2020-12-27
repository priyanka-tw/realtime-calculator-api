package model

//go:generate mockgen -source=client.go -destination=../mock/conn_mock.go -package=mock

type Conn interface {
	WriteJSON(v interface{}) error
	ReadJSON(v interface{}) error
	Close() error
}

type Client struct {
	Connection Conn
	Username string
}
