package utils

import (
	"io"
	"net/http"
)

func HttpGet(url string) (*http.Response, error) {
	return http.Get(url)
}

func ReadResponseBody(response *http.Response) string {
	body, _ := io.ReadAll(response.Body)
	return string(body)
}
