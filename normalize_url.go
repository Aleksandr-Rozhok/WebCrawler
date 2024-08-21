package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawUrl string) (string, error) {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	path := strings.TrimSuffix(parsedURL.Path, "/")
	hostname := strings.TrimPrefix(parsedURL.Hostname(), "www.")
	normalizedURL := hostname + path

	return normalizedURL, nil
}
