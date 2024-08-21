package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedRawBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Ошибка парсинга URL 1: %s : %v", rawBaseURL, err)
		return
	}

	parsedRawCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Ошибка парсинга URL 1: %s : %v", rawCurrentURL, err)
		return
	}

	host1 := parsedRawBaseURL.Hostname()
	host2 := parsedRawCurrentURL.Hostname()

	if host1 != host2 {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
	}

	if pages[normalizedURL] > 0 {
		pages[normalizedURL]++
		return
	} else {
		pages[normalizedURL] = 1

		body, err := getHTML(rawCurrentURL)
		if err != nil {
			fmt.Println(err)
		}

		urls, err := getURLsFromHTML(body, rawBaseURL)
		if err != nil {
			fmt.Println(err)
		}

		for _, nextURL := range urls {
			crawlPage(rawBaseURL, nextURL, pages)
		}
	}
}
