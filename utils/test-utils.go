package utils

import (
	"net/http"
)

func APIClient(url string, token string) *http.Response{
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		return &http.Response{}
	}
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Authorization": {"Bearer Token"},
	}
	client := &http.Client{}
	resp, _ := client.Do(req)
	return resp
}