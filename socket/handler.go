package socket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	model2 "realtime-calculator-api/calculator/model"
	"realtime-calculator-api/socket/model"
)

type Handler struct {
	upgrader    UpgraderWrapper
	hub         Hub
	ehGenerator EventHandlerGenerator
}

func NewSocketHandler(upgrader UpgraderWrapper, hub Hub, ehGenerator EventHandlerGenerator) Handler {
	return Handler{upgrader: upgrader, hub: hub, ehGenerator: ehGenerator}
}

func (wsh Handler) ServeWrapper(ctx *gin.Context) {
	err := wsh.Serve(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.Println("socket handler: ServeWrapper request completed")
	ctx.Status(200)
}

func (wsh Handler) Serve(w http.ResponseWriter, r *http.Request) error {
	log.Println("socket handler: Serve request initiated")
	wsConnection, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("unable to upgrade http to ws, err: %s", err.Error())
		return err
	}

	client := &model.Client{Connection: wsConnection}
	wsh.hub.RegisteredClients()[client] = true

	err = wsh.ListenForEvents(client)
	if err != nil {
		log.Println("error encountered while listening for event, err: ", err)
		return err
	}

	return nil
}

func (wsh Handler) ListenForEvents(currentClient *model.Client) error {
	defer func() {
		delete(wsh.hub.RegisteredClients(), currentClient)
		currentClient.Connection.Close()
	}()

	log.Println("socket handler: listening for events")
	for {
		var ev model.EventMetadata
		err := currentClient.Connection.ReadJSON(&ev)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		err = wsh.triggerEvent(nil, ev)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (wsh Handler) BroadcastResult(ctx *gin.Context) {

	calc, ok := ctx.Get("calculator")
	if !ok {
		log.Println("no data for broadcasting")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	calculator := calc.(model2.Calculator)
	ev := model.EventMetadata{Event: "calculate", Data: calculator.String()}

	err := wsh.triggerEvent(nil, ev)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (wsh Handler) triggerEvent(currentClient *model.Client, ev model.EventMetadata) error {
	handler, err := wsh.ehGenerator.GetHandler(ev.Event)
	if err != nil {
		log.Println(err)
		return err
	}

	err = handler.Handle(currentClient, ev.Data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
