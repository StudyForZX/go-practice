package rpc_example

type Greeter struct {
}

func (p *Greeter) Greet(request string, response *string) error {
	*response = "Hello: " + request
	return nil
}
