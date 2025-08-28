package rpc_example

import (
	"fmt"
	"net/rpc"
)

func Client() error {
	client, err := rpc.Dial("tcp", "localhost:8888")
	if err != nil {
		return fmt.Errorf("dialing error: %s", err.Error())
	}

	var response string
	err = client.Call("Greeter.Greet", "magic", &response)
	if err != nil {
		return fmt.Errorf("call rpc func error: %s", err.Error())
	}

	fmt.Println(response)

	return nil
}
