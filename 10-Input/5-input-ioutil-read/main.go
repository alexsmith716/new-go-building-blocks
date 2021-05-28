package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 'io' provides 'ioutil' package to read text files
// 'ReadFile()' function takes the path to a text file & returns the text content and any read error

// 'os' provides function 'Open()'

// always check for read errors
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	// 'ReadFile()' returns the entire text file content (a byte slice)
	readFile, err := ioutil.ReadFile("file.txt")
	check(err)

	// use 'string()' function to translate byte code into characters
	fmt.Println("ioutil.ReadFile():",string(readFile))

	// 'Open()' returns an 'os.File' type
	// 'os.File' type has 'Read()', 'Seek()' and 'Close()' methods to process file content 
	openFile, err := os.Open("file.txt")
	check(err)
	// will 'defer' to automatically closing the file at the end of the function
	// defer openFile.Close()

	// specify a position to begin reading the opened file 
	specifiedPosition, err := openFile.Seek(94, 0)
	check(err)

	// create a byte type slice for with a length that represents a number of characters
	byteSlice := make([]byte, (len(readFile)-78))

	// 'Read()' takes the 'byteSlice' to know how many text characters to 'read'
	// 'Read()' then returns the number of bytes/characrers it read
	numberOfBytes, err := openFile.Read(byteSlice)
	check(err)

	openFile.Close()

	fmt.Printf("numberOfBytes: %v \n", numberOfBytes)
	fmt.Printf("specifiedPosition: %v \n", specifiedPosition)
	fmt.Printf("Display text from the second sentence: %v \n", string(byteSlice[:numberOfBytes]))
}

//	% go run main.go
//	ioutil.ReadFile(): Every request your application sends to the API needs to identify your application to Google. There are two ways to identify your application: using an OAuth 2.0 token (which also authorizes the request) and/or using the application's API key.
//	numberOfBytes: 149 
//	specifiedPosition: 94 
//	Display text from the second sentence: There are two ways to identify your application: using an OAuth 2.0 token (which also authorizes the request) and/or using the application's API key.
