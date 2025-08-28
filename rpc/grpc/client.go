package grpc_example

import (
	"context"
	"fmt"
	"log"

	"github.com/studyforzx/lt/rpc/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client() error {
	conn, err := grpc.NewClient(
		":8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("无法连接到服务器：%v", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	addRes, err := c.Add(context.Background(), &pb.AddRequest{
		A: 5,
		B: 3,
	})
	if err != nil {
		return fmt.Errorf("加法调用失败：%v", err)
	}

	log.Printf("5 + 3 = %d\n", addRes.Result)

	multiplyRes, err := c.Multiply(context.Background(), &pb.MultiplyRequest{
		A: 5,
		B: 3,
	})
	if err != nil {
		return fmt.Errorf("乘法调用失败： %v", err)
	}

	log.Printf("5 * 3 = %d\n", multiplyRes.Result)

	return nil
}
