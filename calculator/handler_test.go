package calculator

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"realtime-calculator-api/calculator/mock"
	"realtime-calculator-api/calculator/model"
	"testing"
)

type CalculatorHandlerTestSuite struct {
	suite.Suite
	ctrl             *gomock.Controller
	mockService      *mock.MockService
	testContext      *gin.Context
	responseRecorder *httptest.ResponseRecorder
	handler          CalculateHandler
}

func TestCalculatorHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorHandlerTestSuite))
}

func (suite *CalculatorHandlerTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockService = mock.NewMockService(suite.ctrl)
	suite.responseRecorder = httptest.NewRecorder()
	suite.testContext, _ = gin.CreateTestContext(suite.responseRecorder)

	suite.handler = NewCalculatorHandler(suite.mockService)
}

func (suite *CalculatorHandlerTestSuite) Test_ShouldReturnResult_ForGivenExpression() {
	calculator := model.Calculator{Expression: "5*4"}
	requestBytes, _ := json.Marshal(calculator)
	suite.testContext.Request, _= http.NewRequest("POST", "/calculate", bytes.NewBuffer(requestBytes))
	suite.mockService.EXPECT().Calculate("5*4").Return("20", nil).Times(1)

	suite.handler.Calculate(suite.testContext)

	calculator.Result = "20"
	expected, _ := json.Marshal(calculator)
	assert.Equal(suite.T(), 200, suite.responseRecorder.Code)
	assert.Equal(suite.T(), string(expected), suite.responseRecorder.Body.String())
}

func (suite *CalculatorHandlerTestSuite) Test_ShouldReturnBadRequest_ForInvalidJson() {
	suite.testContext.Request, _= http.NewRequest("GET", "/calculate", nil)

	suite.handler.Calculate(suite.testContext)

	assert.Equal(suite.T(), 400, suite.responseRecorder.Code)
}

func (suite *CalculatorHandlerTestSuite) Test_ShouldReturnBadRequest_ForMissingExpression() {
	calculator := model.Calculator{}
	requestBytes, _ := json.Marshal(calculator)
	suite.testContext.Request, _= http.NewRequest("POST", "/calculate", bytes.NewBuffer(requestBytes))

	suite.handler.Calculate(suite.testContext)

	assert.Equal(suite.T(), 400, suite.responseRecorder.Code)
}

func (suite *CalculatorHandlerTestSuite) Test_ShouldReturnInternalServerError_OnErrorFromService() {
	calculator := model.Calculator{Expression: "5*4"}
	requestBytes, _ := json.Marshal(calculator)
	suite.testContext.Request, _= http.NewRequest("POST", "/calculate", bytes.NewBuffer(requestBytes))
	anError := errors.New("err")
	suite.mockService.EXPECT().Calculate("5*4").Return("", anError).Times(1)

	suite.handler.Calculate(suite.testContext)

	assert.Equal(suite.T(), 500, suite.responseRecorder.Code)
}

func (suite *CalculatorHandlerTestSuite) Test_ShouldAddMetadata_OnContext() {
	calculator := model.Calculator{Expression: "5*4"}
	requestBytes, _ := json.Marshal(calculator)
	suite.testContext.Request, _= http.NewRequest("POST", "/calculate", bytes.NewBuffer(requestBytes))
	suite.mockService.EXPECT().Calculate("5*4").Return("20", nil).Times(1)

	suite.handler.Calculate(suite.testContext)
	actual, _ := suite.testContext.Get("calculator")

	calculator.Result = "20"
	assert.Equal(suite.T(), calculator, actual)
}
