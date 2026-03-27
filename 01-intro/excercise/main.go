package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int { // ← map here is fine — it's a return type
	counts := make(map[string]int) // ← make() inside a function — correct

	for _, word := range strings.Fields(s) {
		counts[word]++
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
