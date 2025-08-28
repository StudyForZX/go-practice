package test

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	grpc_example "github.com/studyforzx/lt/rpc/grpc"
)

type GRPCTestSuite struct {
	suite.Suite
}

func TestGRPCTestSuite(t *testing.T) {
	suite.Run(t, new(GRPCTestSuite))
}

func (suite *GRPCTestSuite) SetupSuite() {
	// 在所有测试之前运行的初始化代码
}

func (suite *GRPCTestSuite) TearDownSuite() {
	// 在所有测试完成后运行的清理代码
}

func (suite *GRPCTestSuite) TestRpc() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		err := grpc_example.Server()
		suite.NoError(err)
	}()

	time.Sleep(1 * time.Second)

	wg.Add(1)
	go func() {
		err := grpc_example.Client()
		suite.NoError(err)
		wg.Done()
	}()

	wg.Done()

	wg.Wait()
}
