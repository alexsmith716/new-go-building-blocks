package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 'bufio' (buffered i/o) package implements buffered I/O
// 'bufio' provides `type Scanner`
// `type Scanner` is a memory buffer which has methods to handle data streams
// "Scanner provides a convenient interface for reading data such as a file of newline-delimited lines of text"

// 'strings' provides `func Fields(s string)`
// 'Fields' splits the string 's' around each instance of one or more consecutive white space characters
// 'Fields' returns a slice of substrings of 's' or an empty slice if 's' contains only white space


func main() {

	// user inputs text
	fmt.Print("Enter some text then press 'return': ")

	// input is stored in 'scannerInstance' memory buffer
	scannerInstance := bufio.NewScanner(os.Stdin)

	// 'scannerInstance' 'Scan()' method reads the input 
	scannerInstance.Scan()

	if scannerInstance.Err() != nil {
		fmt.Println(scannerInstance.Err())
	} else {
		// return the input as a single string value
		fmt.Println("Input returned as single string value:", scannerInstance.Text())
	}

	// =========================================

	fmt.Print("Enter some more text then press 'return': ")
	scannerInstance = bufio.NewScanner(os.Stdin)
	scannerInstance.Scan()

	// return the input separated into individual string items (return a slice of substrings)
	words := strings.Fields(scannerInstance.Text())

	fmt.Println("words: ", words)

	// for loop and print [words]
	if scannerInstance.Err() != nil {
		fmt.Println(scannerInstance.Err())
	} else {
		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}
}

//	% go run main.go
//	Enter some text then press 'return': Scanner provides a convenient interface for reading data
//	Input returned as single string value: Scanner provides a convenient interface for reading data
//	Enter some more text then press 'return': Scanner provides a convenient interface for reading data
//	words:  [Scanner provides a convenient interface for reading data]
//	Scanner
//	provides
//	a
//	convenient
//	interface
//	for
//	reading
//	data
