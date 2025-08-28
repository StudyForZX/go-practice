package test

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	rpc_example "github.com/studyforzx/lt/rpc/rpc"
)

type RpcTestSuite struct {
	suite.Suite
}

func TestRpcTestSuite(t *testing.T) {
	suite.Run(t, new(RpcTestSuite))
}

func (suite *RpcTestSuite) SetupSuite() {
	// 在所有测试之前运行的初始化代码
}

func (suite *RpcTestSuite) TearDownSuite() {
	// 在所有测试完成后运行的清理代码
}

func (suite *RpcTestSuite) TestRpc() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		err := rpc_example.Server()
		suite.NoError(err)
	}()

	time.Sleep(1 * time.Second)

	wg.Add(1)
	go func() {
		err := rpc_example.Client()
		suite.NoError(err)
		wg.Done()
	}()

	wg.Done()

	wg.Wait()
}
