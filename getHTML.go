package main

import (
	"errors"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", errors.New(resp.Status)
	}

	if resp.Header.Get("content-type") != "text/html" {
		return "", errors.New(resp.Header.Get("content-type"))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
