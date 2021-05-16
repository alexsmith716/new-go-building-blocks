package main

import (
	"fmt"
)

// custom function types & instances (the `type` keyword)
// any function that accepts functions as args, or returns functions is a "higher-order" function
// passing functions is efficient use of code

// 'higher-order' function (returning named function `div`)
// `twice` <<<< `custFuncTypeAdd` > `addInstance`
func double(twice func(int, int) int) func(int, int) int {
	fmt.Println("Doubled:", twice(6, 2)*2) // (addInstance * 2)
	div := func(a int, b int) int {
		return (a + b) / 2
	}
	return div
}

func main() {

	// `custFuncTypeAdd` signiture: receive 2 interger args and return 1 interger arg
	// will enable creation of instances of custom func type 'custFuncTypeAdd'
	type custFuncTypeAdd func(int, int) int

	// instance 'addInstance' of custom func type 'custFuncTypeAdd' 
	var addInstance custFuncTypeAdd = func(a int, b int) int {
		return a + b
	}

	fmt.Println("Added:", addInstance(6, 2))

	// now use `div` which is based on instance of `custFuncTypeAdd` and modified by `addInstance`
	div := double(addInstance)
	fmt.Println("Divided:", div(6, 2))

	fmt.Printf("addInstance type: %T \n", addInstance)
	fmt.Printf("double type: %T \n", double)
	fmt.Printf("div type: %T \n", div)
}

//	% go run main.go
//	Added: 8
//	Doubled: 16
//	Divided: 4
//	addInstance type: main.custFuncTypeAdd 
//	dub type: func(func(int, int) int) func(int, int) int 
//	div type: func(int, int) int
