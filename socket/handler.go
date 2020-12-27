package socket

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	upgrader UpgraderWrapper
}

func NewSocketHandler(upgrader UpgraderWrapper) Handler {
	return Handler{upgrader: upgrader}
}

func (wsh Handler) ServeWrapper(ctx *gin.Context) {
	err := wsh.serve(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	log.Println("socket handler: ServeWrapper request completed")
	ctx.Status(200)
}

func (wsh Handler) serve(w http.ResponseWriter, r *http.Request) error {
	log.Println("socket handler: Serve request initiated")
	_, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("unable to upgrade http to ws, err: %s", err.Error())
		return err
	}
	return nil
}
