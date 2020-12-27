package calculator

import (
	"fmt"
	"github.com/apaxa-go/eval"
	"log"
)

type calculatorService struct{}

type Service interface {
	Calculate(expression string) (string, error)
}

func NewCalculatorService() Service {
	return calculatorService{}
}

func (c calculatorService) Calculate(expression string) (string, error) {
	log.Println("calculator service: Calculate expression ", expression)

	expr, err := eval.ParseString(expression, "")
	if err != nil {
		log.Println("error encountered while parsing expression, err: ", err)
		return "", err
	}

	r, err := expr.EvalToInterface(nil)
	if err != nil {
		log.Println("error encountered while evaluating expression, err: ", err)
		return "", err
	}

	log.Println("calculator service: Calculation done")
	return fmt.Sprintf("%v", r), nil
}
