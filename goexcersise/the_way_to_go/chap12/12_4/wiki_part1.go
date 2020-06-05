package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type page struct {
	title string
	body  []byte
}

func (p *page) save() {
	f, err := os.OpenFile(p.title, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("got error when open file, error is %s", err)
		return
	}
	w := bufio.NewWriter(f)
	w.WriteString(string(p.body))
	w.Flush()
}

func (p *page) load() {
	buf, err := ioutil.ReadFile(p.title)
	if err != nil {
		fmt.Printf("got error when open file, error is %s", err)
		return
	}
	fmt.Println(string(buf))
}

func main() {
	p := new(page)
	p.title = "title_test.txt"
	p.body = []byte("test")
	p.save()
	p.load()
}
