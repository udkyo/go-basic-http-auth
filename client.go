package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	username = "abc"
	password = "123"
)

func main() {
	call("http://localhost:8234/example", "POST")
}

func call(url, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.SetBasicAuth(username, password)
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	return nil
}
