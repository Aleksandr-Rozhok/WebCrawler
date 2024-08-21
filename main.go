package main

import "fmt"

func main() {
	url, err := normalizeURL("https://blog.boot.dev/path/")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(url)
}
