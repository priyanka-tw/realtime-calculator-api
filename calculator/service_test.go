package calculator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CalculatorServiceTestSuite struct {
	suite.Suite
	service Service
}

func TestCalculatorServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorServiceTestSuite))
}

func (suite *CalculatorServiceTestSuite) SetupTest() {
	suite.service = NewCalculatorService()
}

func (suite *CalculatorServiceTestSuite) Test_ShouldReturnCalculatedResult_ForGivenExpression() {
	expression := "9.6*7"

	result, err := suite.service.Calculate(expression)

	assert.Equal(suite.T(), "67.2", result)
	assert.Nil(suite.T(), err)
}

func (suite *CalculatorServiceTestSuite) Test_ShouldReturnError_ForParsingError() {
	expression := "ff/"

	result, err := suite.service.Calculate(expression)

	assert.Equal(suite.T(), "", result)
	assert.NotNil(suite.T(), err)
}
