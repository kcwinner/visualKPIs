package routers

import (
	"bytes"
	"net/http"
)

func buildGetRequest(URL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func buildPostRequest(URL string, jsonBody []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return req, nil
}
