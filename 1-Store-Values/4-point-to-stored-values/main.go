package main

import (
  "fmt"
)

// point to stored values
// whenever program creates a container to store data, computer allocates space in memory to store the data
// an initialized variable consists of 3 parts (name, value, memory address)
// a pointer variable can store the address of another variable and can access the value stored at that address
// the declaration of a pointer variable requires the data type to be prefixed by "*" to indicate it's a pointer

// `var ptr *int`	// declares the variable will point to an "int"

// the address of another variable can be assigned to the pointer by prefixing the other variable's name with an ampersand "&"

// `ptr = &num` // assigns an address to the pointer variable

// the value at the assigned address can then be accessed by prefixng the pointer name with an asterisk "*"

// `* ptr`  // points to the value at the assigned address
// the "*" operator performs 2 purposes
//	1) as a type descriptor in a declaration
//	2) as a dereferencer when placed before a pointer name

// the pointer changes the original value stored in the "var" to which it points, not a copy of that value
// pointers are everywhere in Go

// var ptr *int  <<<< declare `var ptr` will point to an `int`
// var num int = 200
// ptr = &num    <<<< assign an address to the pointer var
// *ptr          <<<< points to the value at assigned address

func main() {

	// declare and initialize a integer var and a pointer var
  var num int = 20
  var ptr *int = &num // `var` `ptr` will point "*" to an `int` and be assigned the address `&num`

  // display value and memory address of integer variable
  fmt.Printf("num value: %v type: %T \n", num, num)
  fmt.Printf("num address: %v type: %T \n", ptr, ptr) // "pointer" address of `num`

  fmt.Printf("num via pointer: %v type: %T \n", *ptr, *ptr) // "pointer" value of `num`
  fmt.Printf("ptr address: %v type: %T \n", &ptr, &ptr) // address of "pointer"

  *ptr = 100 // change value stored in `num`
  fmt.Printf("new num value: %v type: %T \n", num, num)
}

//  % go run main.go
//  num value: 20 type: int 
//  num address: 0xc000016070 type: *int 
//  num via pointer: 20 type: int 
//  ptr address: 0xc00000e028 type: **int 
//  new num value: 100 type: int
