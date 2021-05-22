package main

import (
	"fmt"
	"strings"
)

// a 'string' is zero or more characters enclosed in double quotes
// empty quotes initialize a variable as an empty 'string' value
// the `nil` zero value is a 'string' literal when enclosed in quotes `"nil"`
// a 'string' is a collection/array of characters -like elements of an array
// apply array characteristics when dealing with 'string' values
// `len()` function will return the number of characters
// `+` operator is a concatenation operator for joining 'string' values together

func main() {

	// intialize 2 string variables and a concatenated string
	string1, string2 := "The Go programming launguage ", "is different from JavaScript."

	// `+` operator is a concatenation operator for joining 'string' values together
	concatString := string1 + string2

	stringCharacters, err := fmt.Printf("concatString: %v \n", concatString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("stringCharacters (value default format): ", stringCharacters)
		fmt.Println("len(concatString):", len(concatString))
	}

	stringArray := []string{"A", "string", "is", "an", "array", "of", "characters"}

	// `Join()` function joins/concatenates all elements of a string array -creates a single string
	// `Join()` takes the name of a string and separator value
	stringsJoin := strings.Join(stringArray, " ")

	fmt.Println(stringsJoin)

	fmt.Printf("stringsJoin[0] (default format -ASCII): %v \n", stringsJoin[0])
	fmt.Printf("stringsJoin[0] (character): %c \n", stringsJoin[0])

	fmt.Printf("stringsJoin[1] (default format -ASCII): %v \n", stringsJoin[1])
	fmt.Printf("stringsJoin[1] (character): %c \n", stringsJoin[1])

	fmt.Printf("stringsJoin[2] (default format -ASCII): %v \n", stringsJoin[2])
	fmt.Printf("stringsJoin[2] (character): %c \n", stringsJoin[2])
}

//	% go run main.go
//	concatString: The Go programming launguage is different from JavaScript. 
//	stringCharacters (value default format):  74
//	len(concatString): 58
//	A string is an array of characters
//	stringsJoin[0] (default format -ASCII): 65 
//	stringsJoin[0] (character): A 
//	stringsJoin[1] (default format -ASCII): 32 
//	stringsJoin[1] (character):   
//	stringsJoin[2] (default format -ASCII): 115 
//	stringsJoin[2] (character): s
