package main

import (
  "fmt"
)

// display variable values

func main() {

  num := 100
  pi := 3.1415926536

  // value of variables displayed with "fmt.Println()" function (for string)
	// desired format to display value specified with "fmt.Println()" function using the suitable "format specifier" and var name

	// `fmt.Printf()` does not add a new line after the output. `\n` is needed to add new line
  // output variable values in varous formats
  fmt.Printf("num: %v type:%T \n", num, num)
  fmt.Printf("num: %v type:%T \n", pi, pi)

  fmt.Printf("%%7d displays %7d \n", num)
  fmt.Printf("%%07d displays %07d \n", num)

  fmt.Printf("Pi is approximately %1.10f \n", pi)
  fmt.Printf("Right-aligned %20.3f rounded pi \n", pi)
  fmt.Printf("Left-aligned %-20.3f rounded pi", pi)
}

//  % go run main.go
//  num: 100 type:int 
//  num: 3.1415926536 type:float64 
//  
//  %7d displays     100 
//  %07d displays 0000100 
//  
//  Pi is approximately 3.1415926536 
//  Right-aligned                3.142 rounded pi 
//  Left-aligned 3.142                rounded pi  

//	"format specifier"
//	%s	A string of characters					"some string"
//	%d	An integer											100
//	%f	A floating point number	100			0.123456
//	%c	A single character							"A"
//	%t	A boolean value									true
//	%p	A machine memory address				0x0022FF36
//	%v	The value in a default format		ANY
//	%T	The data type of the variable		int
