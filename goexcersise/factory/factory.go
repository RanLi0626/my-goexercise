package main

import "fmt"

type Client interface {
	Method() string
}

type client struct{}

func (c *client) Method() string {
	return "client.Method"
}

func NewClient() Client {
	return &client{}
}

func main() {
	c := NewClient()
	fmt.Println(c.Method())
}
