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
