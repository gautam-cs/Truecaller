# Service to get the longest prefix

#Algorithm and data structure

* Create a trie data structures using the input words from given file
* Trie creation will not be on runtime. It will inmemory.
* Total alphabet number size will be 62 (26 small letter, 26 capital letter & 10 integers)
* Create a function which will return the longest prefix based on trie search

## Code Structure
* Place service in service
* Place constants in constants. They should not depend on any other internal app package
* Add more testcase inside the defined constants file
* Place test function in testing

## usage of go routine
* running the test cases for different input concurrently

```bash
#go to project Truecaller
go mod init truecaller/server/app
go mod tidy

# run main function
go run main.go

# run test based on data from constants/LongestPrefixSample.go input
go test -v

```

## Contributors
* [gautam.official.it@gmail.com](mailto:gautam.official.it@gmail.com)