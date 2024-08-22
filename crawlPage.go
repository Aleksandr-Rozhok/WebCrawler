package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()
	parsedRawCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Ошибка парсинга URL 1: %s : %v", rawCurrentURL, err)
		return
	}

	host1 := cfg.baseURL.Hostname()
	host2 := parsedRawCurrentURL.Hostname()

	if host1 != host2 {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
	}

	if cfg.addPageVisit(normalizedURL) {

		body, err := getHTML(rawCurrentURL)
		if err != nil {
			fmt.Println(err)
		}

		urls, err := getURLsFromHTML(body, cfg.baseURL.Scheme+"://"+cfg.baseURL.Host)
		if err != nil {
			fmt.Println(err)
		}

		cfg.concurrencyControl <- struct{}{}
		defer func() {
			<-cfg.concurrencyControl
		}()

		for _, nextURL := range urls {
			cfg.wg.Add(1)
			go cfg.crawlPage(nextURL)
		}

	} else {
		return
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if cfg.pages[normalizedURL] > 0 {
		cfg.pages[normalizedURL]++
		return false
	} else {
		cfg.pages[normalizedURL] = 1
		return true
	}
}
