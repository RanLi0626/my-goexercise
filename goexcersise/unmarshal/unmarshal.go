package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goexcersise/schema"
	"io/ioutil"
	"net/http"
)

func main() {
	reqbody := &schema.SharedLink{}
	/*b, err := proto.Marshal(reqbody)
	if err != nil {
		fmt.Println(err)
	}*/
	b, err := json.Marshal(reqbody)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://ran-dot-partner-dot-fieldbrowser-dummy-data.appspot.com/api/v2/sharedlinks", bytes.NewReader(b))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer ya29.a0Ae4lvC1DDsN2n8KKADXDtWimfe276yXUKZWRMRzYhuxCcFyGkKCjpcOCS_hdgSGOzJVPDqlZBfrdg_1_SoRItnoSmS8nuUwStR4mbAYqmWZULQWJe4gXuCWTjJuOIzItGtPbn4ocOJDRdCGx9IvtXYA86XO4v9CKnaiT")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(respbody))

}
