package testify

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MyTestSuite struct {
	suite.Suite
}

func (suite *MyTestSuite) SetupSuite() {
	// 在所有测试之前运行的初始化代码
}

func (suite *MyTestSuite) TearDownSuite() {
	// 在所有测试完成后运行的清理代码
}

func (suite *MyTestSuite) TestStep1() {
	// 按顺序运行的第一个测试步骤
	suite.T().Log("Test Step 1")
}

func (suite *MyTestSuite) TestStep2() {
	// 按顺序运行的第二个测试步骤
	suite.T().Log("Test Step 2")
}

func (suite *MyTestSuite) TestStep3() {
	// 按顺序运行的第三个测试步骤
	suite.T().Log("Test Step 3")
}

// 可以通过一个主测试函数来按顺序执行测试步骤
func (suite *MyTestSuite) TestSequential() {
	suite.Run("Step1", suite.TestStep1)
	suite.Run("Step2", suite.TestStep2)
	suite.Run("Step3", suite.TestStep3)
}

func TestMyTestSuite(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
}
