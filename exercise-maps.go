package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wcmap := make(map[string]int)
	word_array := strings.Fields(s)
	for _, word := range word_array {
		wcmap[word] = wcmap[word] + 1
	}
	return wcmap
}
func main() {
	wc.Test(WordCount)
}
