package event

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"realtime-calculator-api/socket/mock"
	"testing"
)

type EventHandlerSuite struct {
	suite.Suite
	ctrl         *gomock.Controller
	mockHub      *mock.MockHub
	eventHandler EventHandler
}

func TestEventHandlerSuite(t *testing.T) {
	suite.Run(t, new(EventHandlerSuite))
}

func (suite *EventHandlerSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockHub = mock.NewMockHub(suite.ctrl)
	suite.eventHandler = NewEventHandler(suite.mockHub)
}

func (suite *EventHandlerSuite) Test_ShouldReturnLoginHandler_ForLoginEvent() {
	expected := LoginHandler{
		hub:   suite.mockHub,
		Count: &Count{},
	}

	actual, err := suite.eventHandler.getHandler("login")

	assert.Equal(suite.T(), expected, actual)
	assert.Nil(suite.T(), err)
}

func (suite *EventHandlerSuite) Test_ShouldReturnError_ForInvalidHandlerEvent() {
	actual, err := suite.eventHandler.getHandler("test-event")

	assert.Equal(suite.T(), nil, actual)
	assert.NotNil(suite.T(), err)
}
