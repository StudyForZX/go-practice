package grpc_example

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/studyforzx/lt/rpc/grpc/pb"
	"google.golang.org/grpc"
)

type ServerModel struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *ServerModel) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("收到加法请求：%d + %d", in.A, in.B)
	return &pb.AddResponse{Result: in.A + in.B}, nil
}

func (s *ServerModel) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	log.Printf("收到乘法请求：%d * %d", in.A, in.B)
	return &pb.MultiplyResponse{Result: in.A * in.B}, nil
}

func Server() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return fmt.Errorf("无法监听端口：%v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &ServerModel{})

	log.Println("gRPC服务器启动，端口：8080")

	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("服务器启动失败：%v", err)
	}

	return nil
}
