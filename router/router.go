package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"realtime-calculator-api/calculator"
	"realtime-calculator-api/socket"
)

func InitializeRouter() *gin.Engine {

	engine := gin.Default()
	engine.Use(corsMiddleware)

	calculatorService := calculator.NewCalculatorService()
	calculatorHandler := calculator.NewCalculatorHandler(calculatorService)

	var upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	hub := socket.NewHub()
	eventHandlerGenerator := socket.NewEventHandlerGenerator(hub)
	socketHandler := socket.NewSocketHandler(upgrader, hub, eventHandlerGenerator)

	engine.POST("/calculate-and-broadcast", calculatorHandler.Calculate, socketHandler.BroadcastResult)
	engine.GET("/ws", socketHandler.ServeWrapper)

	return engine
}

func corsMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
}
