package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	defaultURL  = "localhost"
	defaultPort = 8080
)

func main() {

	var restClient = &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := restClient.Get(defaultURL)
	if err != nil {
		// TODO log
	}
	defer response.Body.Close()

	//err := ProcessGetResponse(response)

	fmt.Printf("Response: %s", response.Body)
}
