package socket

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"realtime-calculator-api/socket/mock"
	"realtime-calculator-api/socket/model"
	"testing"
)

type HubSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	mockConn *mock.MockConn
	hub      Hub
}

func TestHubSuite(t *testing.T) {
	suite.Run(t, new(HubSuite))
}

func (suite *HubSuite) SetupTest() {

	suite.ctrl = gomock.NewController(suite.T())
	suite.mockConn = mock.NewMockConn(suite.ctrl)
	suite.hub = NewHub()
}

func (suite *HubSuite) Test_Should_WriteJson_ForAllConnections_InHub() {
	suite.setupConnections(2)
	metadata := model.EventMetadata{
		Event: "login",
		Data:  "test-user",
	}
	suite.mockConn.EXPECT().WriteJSON(metadata).Return(nil).Times(2)

	actual := suite.hub.BroadcastToAllClients(metadata)

	assert.Nil(suite.T(), actual)
}

func (suite *HubSuite) Test_Should_ReturnError_IfErrorEncounteredWhileWriteJson() {
	suite.setupConnections(3)
	anError := errors.New("err")
	metadata := model.EventMetadata{
		Event: "login",
		Data:  "test-user",
	}
	suite.mockConn.EXPECT().WriteJSON(metadata).Return(nil).Times(1)
	suite.mockConn.EXPECT().WriteJSON(metadata).Return(anError).Times(1)
	suite.mockConn.EXPECT().WriteJSON(metadata).Return(nil).Times(1)

	actual := suite.hub.BroadcastToAllClients(metadata)

	assert.Equal(suite.T(), anError, actual)
}

func (suite *HubSuite) Test_Should_ReturnRegisteredClients() {
	connection := &websocket.Conn{}
	client := &model.Client{Connection: connection}
	suite.hub.RegisteredClients()[client] = true
	expected := map[*model.Client]bool{client: true}

	actual := suite.hub.RegisteredClients()

	assert.Equal(suite.T(), expected, actual)
}

func (suite *HubSuite) setupConnections(quantity int) {
	for i := 0; i < quantity; i++ {
		client := &model.Client{Connection: suite.mockConn}
		suite.hub.RegisteredClients()[client] = true
	}
}
