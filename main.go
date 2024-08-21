package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(argsWithoutProg) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %s\n", argsWithoutProg[0])
	}

	html, err := getHTML("https://wikipedia.org")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(html)
	}
}
