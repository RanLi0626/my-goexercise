package main

import (
	"log"
	"net"
	"net/rpc"

	"goexcersise/rpc/internal"
)

// HelloService helloService
type HelloService struct{}

// Hello hello
func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}

// RegisterHelloService registerHelloService
func RegisterHelloService(hsi internal.HelloServiceInterface) {
	rpc.RegisterName(internal.HelloServiceName, hsi)
}
