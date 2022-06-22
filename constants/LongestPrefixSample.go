package constants

type TestSample struct {
	Input    string `json:"input"`
	Expected string `json:"expected"`
}

var SampleTest = []TestSample{
	{
		Input:    "abcdefghijk",
		Expected: "abcdef",
	},
	{
		Input:    "abc",
		Expected: "",
	},
	{
		Input:    "8dLfGZU0Tfrefr",
		Expected: "8dLfGZU0T",
	},
}
