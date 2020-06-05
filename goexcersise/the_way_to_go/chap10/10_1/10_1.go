package main

import (
	"time"
)

type Address struct {
	Add string
}

type VCard struct {
	Name    string
	Date    time.Time
	Pciture []byte
	Addr    *Address
}

func main() {

}
