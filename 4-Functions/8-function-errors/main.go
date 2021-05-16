package main

import (
	"fmt"
)

// handling function errors
// handle errors in functions by returning a separate value
// function signiture must include a 'error' return type (final return value in list of returns)

//	func funcName(paramName type) (returnType, error) {
//		if condition !- true {
//			return -1, errorValue
//		}
//		return returnValue, nil
//	}

func isAPositiveInteger(num int) (int, error) {
	fmt.Println("is", num, "A Positive Integer?")
	if num < 1 {
		// `fmt.Errorf()` function returns an error type
		err := fmt.Errorf("%v is not a positive integer", num)
		// return `-1` in place of expected return
 		return -1, err
	}
	// return expected return value with `nil` zero value ('nil' indicates no error)
	return num, nil
}


func main() {

	// intialize at 2; condition (is `i` >= 2?); final expression (decrement `i` by 1)
	for i := 2; i >= -2; i-- {
		res, err := isAPositiveInteger(i)

		// test error handler
		if err != nil {
			fmt.Println("Failed:", err)
		} else {
			fmt.Println(res, "Passed!")
		}
	}
}

//	% go run main.go
//	2 Passed!
//	1 Passed!
//	Failed: 0 not a positive integer
//	Failed: -1 not a positive integer
//	Failed: -2 not a positive integer
