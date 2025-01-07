package client

import (
	"bytes"
	"io"
	"net/http"
)

var apiUrl string = "http://localhost:8080"
var apiKey string = ""

func SendApiRequest(method string, path string, body []byte, params map[string]string) (string, error) {
	req, err := http.NewRequest(method, apiUrl+path, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	// API key for authentication
	req.Header.Set("x-api-key", apiKey)

	// Add query parameters to URL object
	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
