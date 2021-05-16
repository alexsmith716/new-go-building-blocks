package main

import (
	"fmt"
)

// bitwise operators

// manipulate a compact "bit field" representing a set of boolean "flags"

// REVIEW:
// each byte comprises eight bits that can each contain a 1 or 0 to store a binary number
// that binary number represents a decimal value from 0 to 255

func main() {

	// initialize variable comprising eight bits
	var flags byte = 10 // Binary 1010 (1x8 0x4 1x2 0x1)

	// output all bit flag settings
	fmt.Printf("Flag #1:%v\n", (flags&1) > 0)
	fmt.Printf("Flag #2:%v\n", (flags&2) > 0)
	fmt.Printf("Flag #3:%v\n", (flags&4) > 0)
	fmt.Printf("Flag #4:%v\n", (flags&8) > 0)

	// display the decimal value representing 
	fmt.Printf("Flags Value:\t%08b\t%v\n", flags, flags)

	flags = flags >> 1
	fmt.Printf("New Value:\t%08b\t%v\n", flags, flags)
}

//	operator
//	&       AND
//	|       OR
//	^       XOR/NOT
//	&^      AND NOT
//	<<      shift left 
//	>>      shift right

//	% go run main.go
//	Flag #1:false
//	Flag #2:true
//	Flag #3:false
//	Flag #4:true
//	Flags Value:	00001010	10
//	New Value:	00000101	5
