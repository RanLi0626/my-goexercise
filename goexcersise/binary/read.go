package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	read()
	write()
}

func read() {
	var piVar float64
	var boolVar bool
	//piByte := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	piByte := []byte{24, 45, 68, 84, 251, 33, 9, 64}
	boolByte := []byte{0x00}
	piBuffer := bytes.NewReader(piByte)
	boolBuffer := bytes.NewReader(boolByte)
	binary.Read(piBuffer, binary.LittleEndian, &piVar)
	binary.Read(boolBuffer, binary.LittleEndian, &boolVar)

	fmt.Println("pi:", piVar)
	fmt.Println("bool:", boolVar)
}

func write() {
	pi := math.Pi
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &pi)
	fmt.Println("piBuf:", buf.Bytes())
}
