package main

import (
	"fmt"
	"strings"
)

// 'strings' functions that allow a string to be separated into substrings
// functions treat the string like an array of character elements
// `Fields()` function separates a space-separated sentence into array substrings
// `Split()` function returns an array of substrings existing between the specified separator
// `SplitN()` function: like `Split()` but takes 3rd arg to specify number of substrings to return
// `SplitAfter()` function: like `Split()` but includes the separator in the returned substrings
// `SplitAfterN()` function: like `SplitN()` but includes the separator in the returned substrings

// list the string function type and substrings
func list(strFunc string, substrings []string) {

	fmt.Print(strFunc, ": ")

	//	for _, i := range substrings {
	//		fmt.Print("[", i, "] ")
	//	}
	for i := 0; i < len(substrings); i++ {
		fmt.Print("[", substrings[i], "] ")
	}
	fmt.Print("\n")
}


func main() {

	string1 := "Odit, lorem ipsum dolor sit amet, consectetur adipisicing elit."

	list("Fields", strings.Fields(string1))

	list("Split", strings.Split(string1, ","))
	list("SplitN", strings.SplitN(string1, ",", 2))

	list("SplitAfter", strings.SplitAfter(string1, ","))
	list("SplitAfterN", strings.SplitAfterN(string1, ",", 2))
}

//	% go run main.go
//	Fields: [Odit,] [lorem] [ipsum] [dolor] [sit] [amet,] [consectetur] [adipisicing] [elit.] 
//	Split: [Odit] [ lorem ipsum dolor sit amet] [ consectetur adipisicing elit.] 
//	SplitN: [Odit] [ lorem ipsum dolor sit amet, consectetur adipisicing elit.] 
//	SplitAfter: [Odit,] [ lorem ipsum dolor sit amet,] [ consectetur adipisicing elit.] 
//	SplitAfterN: [Odit,] [ lorem ipsum dolor sit amet, consectetur adipisicing elit.] 
