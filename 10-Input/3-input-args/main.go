package main

import (
	"fmt"
	"os"
)

// command-line arguments can be passed as parameters to program
// the parameters take a space-separated list of arguments "Args"
// 'os' package provides an `Args` slice
// the `Args` slice stores the argument values in individual elements

// the first `os.Args[0]` element is always path to the program


func main() {

	// pass space-separated argument values 'hostname port path headers' as parameters to Go command

	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("`os.Args` passed-in argument %v: %v \n", i, os.Args[i])
	}
}

//	% go run main.go hostname port path headers
//	`os.Args` passed-in argument 1: hostname 
//	`os.Args` passed-in argument 2: port 
//	`os.Args` passed-in argument 3: path 
//	`os.Args` passed-in argument 4: headers
