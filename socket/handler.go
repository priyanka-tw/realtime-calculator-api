package socket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
		return err
	}

	return nil
}

func (wsh Handler) ListenForEvents(currentClient *model.Client) error {
	defer func() {
		delete(wsh.hub.RegisteredClients(), currentClient)
		currentClient.Connection.Close()
	}()

	for {
		var ev model.EventMetadata
		err := currentClient.Connection.ReadJSON(&ev)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		handler, err := wsh.ehGenerator.GetHandler(ev.Event)
		if err != nil {
			log.Println("error encountered while getting an event handler, err: ", err)
			return err
		}

		err = handler.Handle(currentClient, ev.Data)
		if err != nil {
			log.Println("error encountered while handling an event, err: ", err)
			return err
		}
	}
	return nil
}
