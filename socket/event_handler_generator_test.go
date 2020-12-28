package socket

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"realtime-calculator-api/socket/mock"
	"testing"
)

type EventHandlerGeneratorSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	mockHub   *mock.MockHub
	generator EventHandlerGenerator
}

func TestEventHandlerFactorySuite(t *testing.T) {
	suite.Run(t, new(EventHandlerGeneratorSuite))
}

func (suite *EventHandlerGeneratorSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockHub = mock.NewMockHub(suite.ctrl)
	suite.generator = NewEventHandlerGenerator(suite.mockHub)
}

func (suite *EventHandlerGeneratorSuite) Test_ShouldReturnLoginHandler_ForLoginEvent() {
	expected := LoginHandler{
		hub:   suite.mockHub,
		Count: &Count{},
	}

	actual, err := suite.generator.GetHandler("login")

	assert.Equal(suite.T(), expected, actual)
	assert.Nil(suite.T(), err)
}

func (suite *EventHandlerGeneratorSuite) Test_ShouldReturnLogoutHandler_ForLogoutEvent() {
	expected := LogoutHandler{
		hub:   suite.mockHub,
		Count: &Count{},
	}

	actual, err := suite.generator.GetHandler("logout")

	assert.Equal(suite.T(), expected, actual)
	assert.Nil(suite.T(), err)
}

func (suite *EventHandlerGeneratorSuite) Test_ShouldReturnCalculateHandler_ForCalculateEvent() {
	expected := CalculateHandler{
		hub: suite.mockHub,
	}

	actual, err := suite.generator.GetHandler("calculate")

	assert.Equal(suite.T(), expected, actual)
	assert.Nil(suite.T(), err)
}

func (suite *EventHandlerGeneratorSuite) Test_ShouldReturnError_ForInvalidHandlerEvent() {
	actual, err := suite.generator.GetHandler("test-event")

	assert.Equal(suite.T(), nil, actual)
	assert.NotNil(suite.T(), err)
}
