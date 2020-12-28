package socket

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"realtime-calculator-api/socket/mock"
	"realtime-calculator-api/socket/model"
	"testing"
)

type CalculateHandlerSuite struct {
	suite.Suite
	ctrl             *gomock.Controller
	mockConn         *mock.MockConn
	mockHub          *mock.MockHub
	calculateHandler CalculateHandler
}

func TestCalculateHandlerSuite(t *testing.T) {
	suite.Run(t, new(CalculateHandlerSuite))
}

func (suite *CalculateHandlerSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockHub = mock.NewMockHub(suite.ctrl)
	suite.mockConn = mock.NewMockConn(suite.ctrl)
	suite.calculateHandler = NewCalculateHandler(suite.mockHub)
}

func (suite *CalculateHandlerSuite) Test_ShouldBroadcastToAllClients() {
	client := &model.Client{Connection: suite.mockConn, Username: "A"}
	metadata := model.EventMetadata{
		Event: "history",
		Data:  "5*5 = 25",
	}
	suite.mockHub.EXPECT().BroadcastToAllClients(metadata).Return(nil).Times(1)

	err := suite.calculateHandler.Handle(client, "5*5 = 25")

	assert.Nil(suite.T(), err)
}

func (suite *CalculateHandlerSuite) Test_ShouldReturnError_IfErrorEncounteredWhileBroadcasting() {
	client := &model.Client{Connection: suite.mockConn, Username: "A"}
	anError := errors.New("err")
	metadata := model.EventMetadata{
		Event: "history",
		Data:  "5*5 = 25",
	}
	suite.mockHub.EXPECT().BroadcastToAllClients(metadata).Return(anError).Times(1)

	actual := suite.calculateHandler.Handle(client, "5*5 = 25")

	assert.Equal(suite.T(), anError, actual)
}
