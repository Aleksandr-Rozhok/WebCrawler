package main

import (
	"reflect"
	"testing"
)

func TestSortMapByValue(t *testing.T) {
	tests := []struct {
		name     string
		inputMap map[string]int
		expected []keyValue
	}{
		{
			name: "sort map with smallest value",
			inputMap: map[string]int{
				"apple":  5,
				"banana": 2,
				"cherry": 8,
				"date":   6,
			},
			expected: []keyValue{{"cherry", 8}, {"date", 6}, {"apple", 5}, {"banana", 2}},
		},
		{
			name: "sort data with urls",
			inputMap: map[string]int{
				"wagslane.dev":       15,
				"wagslane.dev/about": 14,
				"wagslane.dev/posts/optimize-for-simplicit-first": 1,
				"wagslane.dev/tags":   14,
				"wagslane.dev/tags/a": 3,
				"wagslane.dev/boot":   32,
			},
			expected: []keyValue{
				{"wagslane.dev/boot", 32},
				{"wagslane.dev", 15},
				{"wagslane.dev/about", 14},
				{"wagslane.dev/tags", 14},
				{"wagslane.dev/tags/a", 3},
				{"wagslane.dev/posts/optimize-for-simplicit-first", 1},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortMapByValue(tc.inputMap)
			if reflect.DeepEqual(actual, tc.expected) == false {
				t.Errorf("Test %v - '%s' FAIL: expected '%v', got '%v'", i, tc.name, tc.expected, actual)
			}
		})
	}
}
