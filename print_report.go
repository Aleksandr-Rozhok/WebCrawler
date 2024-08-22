package main

import "fmt"

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	sortedMap := sortMapByValue(pages)

	for _, value := range sortedMap {
		fmt.Printf("Found %d internal links to %s \n", value.value, value.key)
	}
}
