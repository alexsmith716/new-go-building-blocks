package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
// create a temporary file to store data during program execution
// 'io/ioutil' package provides the 'TempFile()' function
// 'TempFile()' function takes 2 args (directory to create the file, a filename pattern)
// directory to create the file can be a 'path' or "" empty string (the system default dir) 
// 'TempFile()' function returns a 'File' type


func check(err error) {
	if err != nil {
		fmt.Printf("Was there a 'PathError'?: %T \n", err)
	}
}

func main() {

	// 'TempFile' automatically opens the file ready to write/read
	// filename pattern includes "*", so a random string ('example.553042871.txt') replaces "*"
	tempFile, err := ioutil.TempFile("", "example.*.txt")
	check(err)
	defer tempFile.Close()

	// the 'Name()' method of returned 'tempFile' references the file path
 	fmt.Printf("Temp file path:%v \n", tempFile.Name())

	numberOfBytesRead, err := tempFile.WriteString("Odit, temporibus reprehenderit dolorum!\n")
	check(err)

	tempFileContent, err := ioutil.ReadFile(tempFile.Name())
	check(err)

	fmt.Printf("numberOfBytesRead: %v \n", numberOfBytesRead)
	fmt.Printf("Read the temp file: %v", string(tempFileContent))

	os.Remove(tempFile.Name())

	// 
	_, err = os.Stat(tempFile.Name())
	// 'Stat' returns the FileInfo structure describing file. If there is an error, it will be of type *PathError
	// the returned 'error' demonstrates that the 'tempFile' has been removed
	check(err)
}

//	% go run main.go
//	Temp file path:/var/folders/v8/vfdvfdvfdvfvfvdfdvfdvfd/T/example.595629245.txt 
//	numberOfBytesRead: 40 
//	Read the temp file: Odit, temporibus reprehenderit dolorum!
//	Was there a 'PathError'?: *fs.PathError
