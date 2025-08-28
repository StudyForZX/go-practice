package rpc_example

import (
	"fmt"
	"net"
	"net/rpc"
)

func Server() error {
	rpc.RegisterName("Greeter", new(Greeter))
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		return fmt.Errorf("ListenTCP error: %s", err.Error())
	}

	fmt.Println("Server is running at localhost:8888 ...")

	conn, err := listener.Accept()
	if err != nil {
		return fmt.Errorf("Accept error: %s", err.Error())
	}

	rpc.ServeConn(conn)

	return nil
}
