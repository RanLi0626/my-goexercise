package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	f, err := os.Open("vcard.gob")
	if err != nil {
		fmt.Printf("failed to open file, error is %s", err)
	}
	defer f.Close()

	//reader := bufio.NewReader(f)
	dec := gob.NewDecoder(f)
	var vc VCard
	err = dec.Decode(&vc)
	if err != nil {
		fmt.Printf("failed to decode vc, error is %s", err)
	}
	fmt.Println(vc)
}
