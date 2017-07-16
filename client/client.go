package main

import (
	"io/ioutil"
	"net/http"
)

// ProcessGetResponse processes the response from a GET req
// Caller must close resp body
func ProcessGetResponse(resp *http.Response) error {
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
