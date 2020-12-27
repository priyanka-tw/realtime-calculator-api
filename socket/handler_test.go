package socket

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
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
	suite.responseRecorder = httptest.NewRecorder()
	suite.testContext, _ = gin.CreateTestContext(suite.responseRecorder)

	suite.handler = NewSocketHandler(suite.mockUpgrader, suite.mockHub)
}

func (suite *SocketHandlerTestSuite) Test_ShouldReturnListen_OnAConnectionUntilErrorEncountered() {
	mockClient := &model.Client{Connection: suite.mockConn}
	registeredClients := map[*model.Client]bool{}
	suite.mockConn.EXPECT().ReadMessage().
		Return(1, []uint8("a message"), nil).
		Return(1, []uint8(""), errors.New(""))
	suite.mockHub.EXPECT().RegisteredClients().Return(registeredClients).Times(1)
	suite.mockConn.EXPECT().Close().Return(nil).Times(1)

	suite.handler.ListenForEvents(mockClient)

	assert.Equal(suite.T(), 200, suite.responseRecorder.Code)
}

func (suite *SocketHandlerTestSuite) Test_ShouldReturn500_IfErrorEncounteredWhileUpgradingProtocol() {
	suite.testContext.Request, _ = http.NewRequest("GET", "/ws", nil)
	suite.mockUpgrader.EXPECT().Upgrade(suite.testContext.Writer, suite.testContext.Request, nil).
		Return(nil, errors.New("error")).Times(1)

	suite.handler.ServeWrapper(suite.testContext)

	assert.Equal(suite.T(), 500, suite.responseRecorder.Code)
}
