package main

import (
	"io/ioutil"
	"net/http"
)

func httpCall(method string, url string) ([]byte, error) {

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	return data, err
}
