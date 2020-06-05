package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type product struct {
	title    string
	price    float64
	quantity int
}

func main() {
	pds := []product{}
	file, err := os.Open("products.txt")
	if err != nil {
		fmt.Printf("got error when open products, error is %s", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		s := strings.Split(string(b), ";")
		price, err := strconv.ParseFloat(s[1], 64)
		if err != nil {
			fmt.Printf("parse float error, error is %s", err)
			return
		}
		quantity, err := strconv.ParseInt(s[2], 10, 0)
		if err != nil {
			fmt.Printf("parse int error, error is %s", err)
			return
		}
		pd := product{s[0], price, int(quantity)}
		pds = append(pds, pd)
	}

	fmt.Println(pds)

}
