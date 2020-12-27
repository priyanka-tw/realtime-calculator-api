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

type LoginHandlerSuite struct {
	suite.Suite
	*Count
	ctrl         *gomock.Controller
	mockConn     *mock.MockConn
	mockHub      *mock.MockHub
	loginHandler LoginHandler
}

func TestLoginHandlerSuite(t *testing.T) {
	suite.Run(t, new(LoginHandlerSuite))
}

func (suite *LoginHandlerSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockHub = mock.NewMockHub(suite.ctrl)
	suite.mockConn = mock.NewMockConn(suite.ctrl)
	suite.Count = &Count{}
	suite.loginHandler = NewLoginHandler(suite.Count, suite.mockHub)
}

func (suite *LoginHandlerSuite) Test_ShouldBroadcastToAllClients() {
	client := &model.Client{Connection: suite.mockConn, Username: "A"}
	metadata := model.EventMetadata{
		Event: "logged in users",
		Data:  "1",
	}
	suite.mockHub.EXPECT().BroadcastToAllClients(metadata).Return(nil).Times(1)

	err := suite.loginHandler.Handle(client, "A")

	assert.Nil(suite.T(), err)
}

func (suite *LoginHandlerSuite) Test_ShouldReturnError_IfErrorEncounteredWhileBroadcasting() {
	client := &model.Client{Connection: suite.mockConn, Username: "A"}
	metadata := model.EventMetadata{
		Event: "logged in users",
		Data:  "1",
	}
	anError := errors.New("err")
	suite.mockHub.EXPECT().BroadcastToAllClients(metadata).Return(anError).Times(1)

	actual := suite.loginHandler.Handle(client, "A")

	assert.Equal(suite.T(), anError, actual)
}

func (suite *LoginHandlerSuite) Test_IncrementCounterOfLoggedInUser_WithEachLoginEvent() {
	client := &model.Client{Connection: suite.mockConn, Username: "A"}
	metadata := model.EventMetadata{
		Event: "logged in users",
		Data:  "3",
	}
	suite.mockHub.EXPECT().BroadcastToAllClients(gomock.Any()).Return(nil).Times(1)
	suite.mockHub.EXPECT().BroadcastToAllClients(gomock.Any()).Return(nil).Times(1)
	suite.mockHub.EXPECT().BroadcastToAllClients(metadata).Return(nil).Times(1)
	_ = suite.loginHandler.Handle(client, "A")
	_ = suite.loginHandler.Handle(client, "A")

	actual := suite.loginHandler.Handle(client, "A")

	assert.Nil(suite.T(), actual)
}
