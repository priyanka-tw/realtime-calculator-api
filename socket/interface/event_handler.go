package _interface


import "realtime-calculator-api/socket/model"

type EventHandler interface {
	Handle(currentClient *model.Client, data string) error
}
