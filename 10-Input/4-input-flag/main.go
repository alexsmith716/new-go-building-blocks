package main

import (
	"fmt"
	"flag"
)

// command-line 'flags' specify options for command-line programs

// 'flag' provides functions for command-line flag parsing

// flag functions: 'String()', 'Bool()', 'Int()', etc.
// each function requires 3 arguments: (name, default value, help message)

// The return value of a flag function is the address of a variable (pointer) that stores the value of the flag
 
func main() {

	// possible to declare an option that uses an existing 'var' declared elsewhere in the program
	// must pass in a pointer to the flag declaration function
	var strVar string
	flag.StringVar(&strVar, "strVar", "foobar", "help message for strVar")

	var intVar int
	flag.IntVar(&intVar, "intVar", 1234, "help message for intVar")

	generateCpuProfileVar := flag.String("generateProfile", "profile.profile", "Emit a profile for debugging")
	maxNodeModuleJsDepthVar := flag.Int("maxDepth", 0, "Specify the maximum depth")
	allowJsVar := flag.Bool("allowFiles", false, "Allow files to be a part of your program")

	// parse the command-line flags
	// called after all flags are defined and before flags are accessed by the program
	flag.Parse()

	fmt.Println("strVar:", strVar)
	fmt.Println("intVar:", intVar)

	// >>>> "*" <<<<< referencing the VALUE stored at the address, NOT the variable address
	fmt.Println("generateProfile:", *generateCpuProfileVar)
	fmt.Println("maxDepth:", *maxNodeModuleJsDepthVar)
	fmt.Println("allowFiles:", *allowJsVar)
	fmt.Println("tail:", flag.Args()) // output any remaining input
}

//	% go run main.go -strVar=stringFlag -intVar=99999999999 a1 a2 a3
//	strVar: stringFlag
//	intVar: 99999999999
//	generateProfile: profile.profile
//	maxDepth: 0
//	allowFiles: false
//	tail: [a1 a2 a3]
