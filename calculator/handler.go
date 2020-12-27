package calculator

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"realtime-calculator-api/calculator/model"
)

type CalculateHandler struct {
	calService Service
}

func NewCalculatorHandler(calculatorService Service) CalculateHandler {
	return CalculateHandler{calService: calculatorService}
}

func (c CalculateHandler) Calculate(context *gin.Context) {
	log.Println("calculate handler: Calculate, request initiated")
	var calculator model.Calculator
	err := context.ShouldBindJSON(&calculator)
	if err != nil {
		log.Println("unable to bind with request object")
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if calculator.Expression == "" {
		log.Println("missing mandatory request parameter")
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := c.calService.Calculate(calculator.Expression)
	if err != nil {
		log.Println("error received from service, err: ", err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	calculator.Result = result
	context.Set("calculator", calculator)

	log.Println("calculate handler: Calculate, request completed")
	context.JSON(http.StatusOK, calculator)
}
