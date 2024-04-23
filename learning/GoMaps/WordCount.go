package main

import (
	"strings"
)

// should return the amount of times each word occurs
func WordCount(s string) map[string]int {
	var WordCounts = make(map[string]int)

	for _, Word := range strings.Fields(s) {
		WordCounts[Word]++
	}

	return WordCounts
}
