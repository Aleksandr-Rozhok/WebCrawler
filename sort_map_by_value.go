package main

import (
	"sort"
)

type keyValue struct {
	key   string
	value int
}

func sortMapByValue(m map[string]int) []keyValue {
	sliceOfPairs := make([]keyValue, 0, len(m))
	for k, v := range m {
		sliceOfPairs = append(sliceOfPairs, keyValue{k, v})
	}

	sort.Slice(sliceOfPairs, func(i, j int) bool {
		return sliceOfPairs[i].value > sliceOfPairs[j].value
	})

	return sliceOfPairs
}
