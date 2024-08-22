package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(argsWithoutProg) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %s\n", argsWithoutProg[0])
	}

	parsedBaseUrl, err := url.Parse(argsWithoutProg[0])
	if err != nil {
		fmt.Printf("error parsing base url: %s\n", err)
	}

	maxConcurrencyArg, err := strconv.Atoi(argsWithoutProg[1])
	if err != nil {
		fmt.Printf("error parsing max concurrency: %s\n", err)
	}
	maxPagesArg, err := strconv.Atoi(argsWithoutProg[2])
	if err != nil {
		fmt.Printf("error parsing max pages: %s\n", err)
	}

	cfg := config{
		pages:              map[string]int{},
		baseURL:            parsedBaseUrl,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrencyArg),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPagesArg,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(argsWithoutProg[0])

	cfg.wg.Wait()
	fmt.Println("crawling finished")
	printReport(cfg.pages, argsWithoutProg[0])
}
