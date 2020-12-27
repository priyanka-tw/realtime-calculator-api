package router

import (
	"github.com/gin-gonic/gin"
	"realtime-calculator-api/calculator"
)

func InitializeRouter() *gin.Engine {
	engine := gin.Default()

	calculatorService := calculator.NewCalculatorService()
	calculatorHandler := calculator.NewCalculatorHandler(calculatorService)

	engine.POST("/calculate", calculatorHandler.Calculate)

	return engine
}
