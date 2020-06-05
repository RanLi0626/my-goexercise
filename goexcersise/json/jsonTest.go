package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type response struct {
	Error err `json:"error"`
}

type err struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Name    string `json:"name"`
}

const (
	link = "https://www.zhihu.com/api/v4/answer_later/count"
)

func getResult(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return res.Body, nil
}

func parseJSON(result io.Reader) (response, error) {
	var resp response
	err := json.NewDecoder(result).Decode(&resp)
	if err != nil {
		return response{}, err
	}
	return resp, nil
}

func storeToFile(resp response) {
	rs, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("jsonFile.json", rs, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	reader, err := getResult(link)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := parseJSON(reader)
	if err != nil {
		log.Fatal(err)
	}
	storeToFile(resp)
}
