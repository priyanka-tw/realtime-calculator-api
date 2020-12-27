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

type LogoutHandlerSuite struct {
	suite.Suite
	*Count
	ctrl          *gomock.Controller
	mockConn      *mock.MockConn
	mockHub       *mock.MockHub
	logoutHandler LogoutHandler
}

func TestLogoutHandlerSuite(t *testing.T) {
	suite.Run(t, new(LogoutHandlerSuite))
}

func (suite *LogoutHandlerSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockHub = mock.NewMockHub(suite.ctrl)
	suite.mockConn = mock.NewMockConn(suite.ctrl)
	suite.Count = &Count{numberOfUsers: 10}
	suite.logoutHandler = NewLogoutHandler(suite.Count, suite.mockHub)
}

func (suite *LogoutHandlerSuite) Test_ShouldBroadcastToAllClients() {
	client := &model.Client{Connection: suite.mockConn, Username: "A"}
	registeredClients := map[*model.Client]bool{client: true}
	metadata := model.EventMetadata{
		Event: "logged in users",
		Data:  "9",
	}
	suite.mockHub.EXPECT().RegisteredClients().Return(registeredClients).Times(1)
	suite.mockHub.EXPECT().BroadcastToAllClients(metadata).Return(nil).Times(1)

	err := suite.logoutHandler.Handle(client, "A")

	assert.Nil(suite.T(), err)
}

func (suite *LogoutHandlerSuite) Test_ShouldReturnError_IfErrorEncounteredWhileBroadcasting() {
	client := &model.Client{Connection: suite.mockConn, Username: "A"}
	registeredClients := map[*model.Client]bool{client: true}
	anError := errors.New("err")
	metadata := model.EventMetadata{
		Event: "logged in users",
		Data:  "9",
	}
	suite.mockHub.EXPECT().RegisteredClients().Return(registeredClients).Times(1)
	suite.mockHub.EXPECT().BroadcastToAllClients(metadata).Return(anError).Times(1)

	actual := suite.logoutHandler.Handle(client, "A")

	assert.Equal(suite.T(), anError, actual)
}

func (suite *LogoutHandlerSuite) Test_DecrementCounterOfLoggedInUser_WithEachLogoutEvent() {
	clientA := &model.Client{Connection: suite.mockConn, Username: "A"}
	clientB := &model.Client{Connection: suite.mockConn, Username: "B"}
	clientC := &model.Client{Connection: suite.mockConn, Username: "C"}
	clientD := &model.Client{Connection: suite.mockConn, Username: "D"}
	registeredClients := map[*model.Client]bool{clientA: true, clientB: true, clientC: true, clientD: true}
	metadata := model.EventMetadata{
		Event: "logged in users",
		Data:  "7",
	}
	suite.mockHub.EXPECT().RegisteredClients().Return(registeredClients).Times(1)
	suite.mockHub.EXPECT().RegisteredClients().Return(registeredClients).Times(1)
	suite.mockHub.EXPECT().RegisteredClients().Return(registeredClients).Times(1)
	suite.mockHub.EXPECT().BroadcastToAllClients(gomock.Any()).Return(nil).Times(1)
	suite.mockHub.EXPECT().BroadcastToAllClients(gomock.Any()).Return(nil).Times(1)
	suite.mockHub.EXPECT().BroadcastToAllClients(metadata).Return(nil).Times(1)
	_ = suite.logoutHandler.Handle(clientA, "A")
	_ = suite.logoutHandler.Handle(clientB, "B")

	actual := suite.logoutHandler.Handle(clientC, "C")

	assert.Nil(suite.T(), actual)
}
