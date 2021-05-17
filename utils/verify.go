package utils

import (
	"fmt"
)


func IsAPositiveInteger(num int) (int, error) {

	if num < 1 {
		err := fmt.Errorf("%v not a positive integer", num)
		return -1, err
	}
	return num, nil
}

//	% go run main.go
//	2 Passed!
//	1 Passed!
//	Failed: 0 not a positive integer
//	Failed: -1 not a positive integer
//	Failed: -2 not a positive integer
