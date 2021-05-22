package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://golang.org/pkg/strings
// 'strings' functions convert the character case of strings
// `ToUpper()` returns all characters uppercase
// `ToLower()` returns all characters lowercase
// `Title()` returns 1st character of each word as uppercase -string arg has to be `ToLower()` 
// `Trim()` removes space from beginning and end of a string


func main() {

	string1 := "Odit, LOREM ipsum Dolor sit amet, consectetur Adipisicing elit."

	fmt.Println(strings.ToUpper(string1))
	fmt.Println(strings.ToLower(string1))
	fmt.Println(strings.Title(strings.ToLower(string1)))

	string1 = "    821    " // 8 bank spaces
	fmt.Printf("%v Type: %T, Length: %v \n", string1, string1, len(string1))

	string1 = strings.Trim(string1, " ")
	fmt.Printf("%v Type: %T, Length: %v \n", string1, string1, len(string1))

	// "strconv.Atoi()" function converts a string into an "int" data type
	num, err := strconv.Atoi(string1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v Type: %T \n", num, num)
	}
}

//	% go run main.go
//	ODIT, LOREM IPSUM DOLOR SIT AMET, CONSECTETUR ADIPISICING ELIT.
//	odit, lorem ipsum dolor sit amet, consectetur adipisicing elit.
//	Odit, Lorem Ipsum Dolor Sit Amet, Consectetur Adipisicing Elit.
//	    821     Type: string, Length: 11 
//	821 Type: string, Length: 3 
//	821 Type: int
