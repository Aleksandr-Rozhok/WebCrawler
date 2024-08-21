package main

import (
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	result := make([]string, 0)
	htmlReader := strings.NewReader(htmlBody)
	htmlTree, err := html.Parse(htmlReader)
	if err != nil {
		return []string{}, err
	}

	traverseNodes(htmlTree, 0, &result)

	for i, link := range result {
		u, err := url.Parse(link)
		if err != nil || !u.IsAbs() {
			result[i] = strings.TrimSuffix(rawBaseURL, "/") + link
		}
	}
	return result, nil
}

func traverseNodes(node *html.Node, depth int, result *[]string) {
	if node == nil {
		return
	}

	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}

	if node.Type == html.ElementNode {
		if node.Data == "a" {
			for _, value := range node.Attr {
				if value.Key == "href" {
					*result = append(*result, value.Val)
				}
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		traverseNodes(child, depth+1, result)
	}
}
