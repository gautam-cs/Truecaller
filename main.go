package main

import (
	"fmt"
	"truecaller/server/app/service"
)

func main() {
	wordsToFind := []string{"abcdefghijk", "abc"}
	trie := service.CreateTrie()
	fmt.Println()
	for i := 0; i < len(wordsToFind); i++ {
		maximumMatchedPrefix := trie.LongestMatchedPrefix(wordsToFind[i])
		result := ""
		if maximumMatchedPrefix == "" {
			result = fmt.Sprintf("no match as all prefixes are greater than input string.")
		} else {
			result = fmt.Sprintf("%s as it is the longest matched prefix.", maximumMatchedPrefix)
		}
		fmt.Println(result)
	}
}
