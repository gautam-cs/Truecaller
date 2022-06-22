package testing

import (
	"testing"
	"truecaller/server/app/constants"
	"truecaller/server/app/service"
)

func TestAdd(t *testing.T) {
	trie := service.CreateTrie()
	sampleTest := constants.SampleTest
	for _, test := range sampleTest {
		validateTest(trie, test, t)
	}
}

func validateTest(trie service.Trie, test constants.TestSample, t *testing.T) {
	output := service.Trie.LongestMatchedPrefix(trie, test.Input)
	if output != test.Expected && output == "" {
		t.Error("failed")
	} else if output == test.Expected && output == "" {
		t.Log("as it is the longest matched prefix.")
	} else {
		t.Logf("%s as it is the longest matched prefix.", output)
	}
}
