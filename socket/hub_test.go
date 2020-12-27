package socket

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"realtime-calculator-api/socket/model"
	"testing"
)

type HubSuite struct {
	suite.Suite
	hub      Hub
}

func TestHubSuite(t *testing.T) {
	suite.Run(t, new(HubSuite))
}

func (suite *HubSuite) SetupTest() {
	suite.hub = NewHub()
}

func (suite *HubSuite) Test_Should_ReturnRegisteredClients() {
	connection := &websocket.Conn{}
	client := &model.Client{Connection: connection}
	suite.hub.RegisteredClients()[client] = true
	expected := map[*model.Client]bool{client: true}

	actual := suite.hub.RegisteredClients()

	assert.Equal(suite.T(), expected, actual)
}

