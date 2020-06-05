package main

import (
	"fmt"
	"log"

	"goexcersise/rpc/internal"
)

func main() {
	client, err := internal.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
