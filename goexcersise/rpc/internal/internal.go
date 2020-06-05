package internal

import "net/rpc"

// HelloServiceName helloServiceName
const HelloServiceName = "path/to/pkg.HelloService"

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

// HelloServiceInterface helloServiceInterface
type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

// HelloServiceClient helloServiceClient
type HelloServiceClient struct {
	*rpc.Client
}

// Hello hello
func (h *HelloServiceClient) Hello(request string, reply *string) error {
	return h.Client.Call(HelloServiceName+".Hello", request, reply)
}

// DialHelloService dialHelloService
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: client}, nil
}
