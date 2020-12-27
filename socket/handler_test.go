package socket

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	model2 "realtime-calculator-api/calculator/model"
	mock2 "realtime-calculator-api/socket/interface/mock"
	"realtime-calculator-api/socket/mock"
	"realtime-calculator-api/socket/model"
	"testing"
)

type SocketHandlerTestSuite struct {
	suite.Suite
	ctrl             *gomock.Controller
	mockUpgrader     *mock.MockUpWrapper
	mockHub          *mock.MockHub
	mockConn         *mock.MockConn
	mockGenerator    *mock.MockEventHandlerGenerator
	mockEventHandler *mock2.MockEventHandler
	testContext      *gin.Context
	responseRecorder *httptest.ResponseRecorder
	handler          Handler
}

func TestSocketHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(SocketHandlerTestSuite))
}

func (suite *SocketHandlerTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockUpgrader = mock.NewMockUpWrapper(suite.ctrl)
	suite.mockHub = mock.NewMockHub(suite.ctrl)
	suite.mockConn = mock.NewMockConn(suite.ctrl)
	suite.mockGenerator = mock.NewMockEventHandlerGenerator(suite.ctrl)
	suite.mockEventHandler = mock2.NewMockEventHandler(suite.ctrl)
	suite.responseRecorder = httptest.NewRecorder()
	suite.testContext, _ = gin.CreateTestContext(suite.responseRecorder)

	suite.handler = NewSocketHandler(suite.mockUpgrader, suite.mockHub, suite.mockGenerator)
}

func (suite *SocketHandlerTestSuite) Test_ShouldListen_OnAConnectionUntilErrorEncountered() {
	mockClient := &model.Client{Connection: suite.mockConn}
	registeredClients := map[*model.Client]bool{}
	suite.mockConn.EXPECT().ReadJSON(gomock.Any()).
		Return(nil).
		Return(errors.New("an error"))
	suite.mockHub.EXPECT().RegisteredClients().Return(registeredClients).Times(1)
	suite.mockConn.EXPECT().Close().Return(nil).Times(1)

	actual := suite.handler.ListenForEvents(mockClient)

	assert.Nil(suite.T(), actual)
}

func (suite *SocketHandlerTestSuite) Test_ShouldListenAndHandleEvent_OnAConnectionUntilErrorEncountered() {
	mockClient := &model.Client{Connection: suite.mockConn}
	registeredClients := map[*model.Client]bool{}
	metadata := &model.EventMetadata{}
	suite.mockConn.EXPECT().ReadJSON(metadata).
		Return(nil)
	suite.mockGenerator.EXPECT().GetHandler(gomock.Any()).Return(suite.mockEventHandler, nil).Times(1)
	suite.mockEventHandler.EXPECT().Handle(mockClient, gomock.Any()).Return(nil).Times(1)
	suite.mockConn.EXPECT().ReadJSON(metadata).
		Return(errors.New("an error"))
	suite.mockHub.EXPECT().RegisteredClients().Return(registeredClients).Times(1)
	suite.mockConn.EXPECT().Close().Return(nil).Times(1)

	actual := suite.handler.ListenForEvents(mockClient)

	assert.Nil(suite.T(), actual)
}

func (suite *SocketHandlerTestSuite) Test_ShouldReturn_IfErrorEncounteredWhileGettingEventHandler() {
	mockClient := &model.Client{Connection: suite.mockConn}
	registeredClients := map[*model.Client]bool{}
	metadata := &model.EventMetadata{}
	suite.mockConn.EXPECT().ReadJSON(metadata).
		Return(nil)
	suite.mockGenerator.EXPECT().GetHandler(gomock.Any()).Return(suite.mockEventHandler, errors.New("an error")).Times(1)
	suite.mockHub.EXPECT().RegisteredClients().Return(registeredClients).Times(1)
	suite.mockConn.EXPECT().Close().Return(nil).Times(1)

	actual := suite.handler.ListenForEvents(mockClient)

	assert.NotNil(suite.T(), actual)
}

func (suite *SocketHandlerTestSuite) Test_ShouldReturn_IfErrorEncounteredWhileHandlingAnEvent() {
	mockClient := &model.Client{Connection: suite.mockConn}
	registeredClients := map[*model.Client]bool{}
	metadata := &model.EventMetadata{}
	suite.mockConn.EXPECT().ReadJSON(metadata).
		Return(nil)
	suite.mockGenerator.EXPECT().GetHandler(gomock.Any()).Return(suite.mockEventHandler, nil).Times(1)
	suite.mockEventHandler.EXPECT().Handle(mockClient, gomock.Any()).Return(errors.New("an error")).Times(1)
	suite.mockHub.EXPECT().RegisteredClients().Return(registeredClients).Times(1)
	suite.mockConn.EXPECT().Close().Return(nil).Times(1)

	actual := suite.handler.ListenForEvents(mockClient)

	assert.NotNil(suite.T(), actual)
}

func (suite *SocketHandlerTestSuite) Test_ShouldTriggerLogoutEvent_IfConnectedUserIsLoggedIn() {
	client := &model.Client{
		Connection: suite.mockConn,
		Username:   "test-username",
	}
	suite.mockConn.EXPECT().ReadJSON(gomock.Any()).Return(errors.New("an error")).Times(1)
	suite.mockHub.EXPECT().RegisteredClients().Return(map[*model.Client]bool{}).Times(1)
	suite.mockGenerator.EXPECT().GetHandler("logout").Return(suite.mockEventHandler, nil).Times(1)
	suite.mockEventHandler.EXPECT().Handle(client, gomock.Any()).Return(nil).Times(1)
	suite.mockConn.EXPECT().Close().Return(nil).Times(1)

	err := suite.handler.ListenForEvents(client)

	suite.NotNil(suite.T(), err)
}

func (suite *SocketHandlerTestSuite) Test_ShouldReturn500_IfErrorEncounteredWhileUpgradingProtocol() {
	suite.testContext.Request, _ = http.NewRequest("GET", "/ws", nil)
	suite.mockUpgrader.EXPECT().Upgrade(suite.testContext.Writer, suite.testContext.Request, nil).
		Return(nil, errors.New("error")).Times(1)

	suite.handler.ServeWrapper(suite.testContext)

	assert.Equal(suite.T(), 500, suite.responseRecorder.Code)
}

func (suite *SocketHandlerTestSuite) Test_ShouldReturn200_OnBroadcastingDataToAllClients() {
	calc := model2.Calculator{
		Expression: "5+4",
		Result:     "9",
	}
	suite.testContext.Set("calculator", calc)
	suite.mockGenerator.EXPECT().GetHandler(gomock.Any()).Return(suite.mockEventHandler, nil).Times(1)
	suite.mockEventHandler.EXPECT().Handle(nil, gomock.Any()).Return(nil).Times(1)

	suite.handler.BroadcastResult(suite.testContext)

	assert.Equal(suite.T(), 200, suite.responseRecorder.Code)
}

func (suite *SocketHandlerTestSuite) Test_ShouldReturn400_IfNoDataForBroadcasting() {

	suite.handler.BroadcastResult(suite.testContext)

	assert.Equal(suite.T(), 400, suite.responseRecorder.Code)
}
