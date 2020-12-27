package router

import (
	"github.com/gin-gonic/gin"
	"realtime-calculator-api/calculator"
)

func InitializeRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(corsMiddleware)

	calculatorService := calculator.NewCalculatorService()
	calculatorHandler := calculator.NewCalculatorHandler(calculatorService)

	engine.POST("/calculate", calculatorHandler.Calculate)

	return engine
}

func corsMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
}