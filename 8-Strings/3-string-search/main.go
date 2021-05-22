package main

import (
	"fmt"
	"strings"
)

// 'strings' functions that search a string to find specified character or substring
// `Contains()` function returns boolean if specified pattern exists in a string
// `index()` function returns integer of first matching pattern in a string (-1 if fails)
// `Count()` function returns integer of a matching patterns in a string (-1 if fails)
// `HasPrefix()` function returns boolean if specified pattern exists as element 'Prefix' 
// `HasSuffix()` function returns boolean if specified pattern exists as element 'Suffix' 

func main() {

	string1 := "Odit, lorem ipsum dolor sit amet, consectetur adipisicing elit."

	fmt.Printf("strings.Contains? 'elit': %v \n", strings.Contains(string1, "elit"))
	fmt.Printf("strings.Contains? 'elitZ': %v \n", strings.Contains(string1, "elitZ"))

	fmt.Printf("strings.Index? 'elit': %v \n", strings.Index(string1, "elit"))

	fmt.Printf("strings.Count? 'p': %v \n", strings.Count(string1, "p"))

	fmt.Printf("strings.HasPrefix? 'lit.': %v \n", strings.HasPrefix(string1, "lit."))
	fmt.Printf("strings.HasSuffix? 'lit.': %v \n", strings.HasSuffix(string1, "lit."))

	fmt.Println(strings.Replace(string1, "adipisicing", "REPLACED!", 1))
}

//	% go run main.go
//	strings.Contains? 'elit': true 
//	strings.Contains? 'elitZ': false 
//	strings.Index? 'elit': 58 
//	strings.Count? 'p': 2 
//	strings.HasPrefix? 'lit.': false 
//	strings.HasSuffix? 'lit.': true 
//	Odit, lorem ipsum dolor sit amet, consectetur REPLACED! elit.
