package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := base64.RawURLEncoding.EncodeToString([]byte("test"))
	fmt.Println(str)
}
