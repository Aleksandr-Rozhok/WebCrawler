package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode >= 400 {
		return "", errors.New(resp.Status)
	}

	if resp.Header.Get("content-type") != "text/html; charset=utf-8" {
		return "", errors.New(resp.Header.Get("content-type"))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
