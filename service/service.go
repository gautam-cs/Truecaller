package service

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	IntAlphaSize = 62
)

type trieNode struct {
	children  [IntAlphaSize]*trieNode
	isWordEnd bool
}

type Trie struct {
	root *trieNode
}

func InitTrie() *Trie {
	return &Trie{
		root: &trieNode{},
	}
}

func (t *Trie) insert(word string) {
	wordLength := len(word)
	current := t.root
	index := 0
	for i := 0; i < wordLength; i++ {
		if word[i] >= 48 && word[i] <= 57 {
			index = int(word[i] - '0')
		} else if word[i] >= 97 && word[i] <= 122 {
			index = 10 + int(word[i]-'a')
		} else {
			index = 36 + int(word[i]-'A')
		}
		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}
		current = current.children[index]

	}
	current.isWordEnd = true
}

func (t Trie) LongestPrefix(word string) string {
	wordLength := len(word)
	current := t.root
	index := 0

	maximumPrefix := ""
	for i := 0; i < wordLength; i++ {
		temp := word[i]
		if word[i] >= 48 && word[i] <= 57 {
			index = int(word[i] - '0')
		} else if word[i] >= 97 && word[i] <= 122 {
			index = 10 + int(word[i]-'a')
		} else {
			index = 36 + int(word[i]-'A')
		}
		if current.children[index] == nil {
			return maximumPrefix
		}
		maximumPrefix += string(rune(temp))
		current = current.children[index]
	}
	return maximumPrefix
}

func (t Trie) LongestMatchedPrefix(word string) string {
	wordLength := len(word)
	current := t.root
	index := 0

	maximumPrefix := ""
	maximumMatchedPrefix := ""
	for i := 0; i < wordLength; i++ {
		temp := word[i]
		if word[i] >= 48 && word[i] <= 57 {
			index = int(word[i] - '0')
		} else if word[i] >= 97 && word[i] <= 122 {
			index = 10 + int(word[i]-'a')
		} else {
			index = 36 + int(word[i]-'A')
		}
		if current.children[index] == nil {
			return maximumMatchedPrefix
		}
		maximumPrefix += string(rune(temp))
		current = current.children[index]
		if current.isWordEnd == true {
			maximumMatchedPrefix = maximumPrefix
		}
	}
	return maximumMatchedPrefix
}

func CreateTrie() Trie {
	wordsToInsert := make([]string, 0)
	trie := InitTrie()
	projectPath, _ := os.Getwd()
	fmt.Printf(projectPath)
	f, err := os.Open(fmt.Sprintf("%s/service/sample_prefixes.txt", projectPath))
	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		path, err := r.ReadString(10) // 0x0A separator = newline
		if err == io.EOF {
			break
		} else if err == nil {
			wordsToInsert = append(wordsToInsert, strings.TrimSuffix(path, "\n"))
		} else {
			fmt.Println("error reading next line ", err)
		}
	}
	for i := 0; i < len(wordsToInsert); i++ {
		trie.insert(wordsToInsert[i])
	}
	return *trie
}
