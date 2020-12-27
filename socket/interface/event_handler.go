package _interface

//go:generate mockgen -source=event_handler.go -destination=./mock/event_handler_mock.go

import "realtime-calculator-api/socket/model"

type EventHandler interface {
	Handle(currentClient *model.Client, data string) error
}
