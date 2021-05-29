package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 'io' provides 'ioutil' package to read & write text files
// 'WriteFile()' function takes 3 args (path to the file, byte slice of text chars to write, file permission)
// 'WriteFile()' returns an error of fail or a success of nil

// 'io' also provides 'OpenFile()' function to write text files
// 'OpenFile()' takes 3 args (path to the file, flag specifying operation typoe, file permission)

// 'os' flags:
// 'O_RDONLY': open the file read-only
// 'O_WRONLY': open the file write-only
// 'O_APPEND': append data to the file when writing
// 'O_RDWR': open the file read-write
// 'O_CREATE': create a new file if none exists

// https://golang.org/pkg/os/
// https://golang.org/pkg/syscall/ // 'syscall' contains an interface to the low-level operating system primitives

// Constants: Flags to 'OpenFile' wrapping those of the underlying system

// check for errors
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	// the byte slice of text to be written
	txt := []byte("\nLorem ipsum dolor sit amet, consectetur adipisicing elit.\n")

	// write 'txt' to file "file.txt" with permission '0644' and return a 'File' type 
	err := ioutil.WriteFile("file.txt", txt, 0644)
	check(err)

	// open "file.txt" and append data with permission
	file, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_APPEND, 0644)
	check(err)
	defer file.Close()

	// create a new byte slice
	slice := []byte("Odit, temporibus reprehenderit dolorum!\n")
	// the 'File' type method 'Write()' writes the text and returns the bytes read/error
	numberOfBytesRead, err := file.Write(slice)
	check(err)

	fmt.Printf("numberOfBytesRead: %v \n", numberOfBytesRead)
	fmt.Printf("Appended data to the file: %v", string(slice[:numberOfBytesRead]))

	readFile, err := ioutil.ReadFile("file.txt")
	check(err)

	fmt.Println("ReadFile('file.txt'):",string(readFile))
}

//	% go run main.go
//	numberOfBytesRead: 40 
//	Appended data to the file: Odit, temporibus reprehenderit dolorum!
//	ReadFile('file.txt'): 
//	Lorem ipsum dolor sit amet, consectetur adipisicing elit.
//	Odit, temporibus reprehenderit dolorum!
//	
