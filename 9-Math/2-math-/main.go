package main

import (
	"fmt"
	"math"
)

// 'math' functions used to evaluate 'float64' arguments
// functions:
// math.Abs(), math.Acos(), math.Asin(), math.Atan(), math.Atan2(), math.Ceil(), math.Cos(),
// math.Exp(), math.Floor(), math.isNaN(), math.Log(), math.Max(), math.Min(), math.Pow(),
// math.Round(), math.Sin(), math.Sqrt(), math.Tan()

func main() {

	// intialize two variables with power values using function `math.Pow()`
	// returns value of 6 to the power of 3 (6*6*6)
	sixToTheThird := math.Pow(6, 3)
	// returns value of 2 to the power of 4 (2*2*2*2)
	twoToTheFourth := math.Pow(2, 4)

	// compare the positive value of the variables and show the largest & smallest numbers
	// return the larger (`math.Max()`) & smaller (`math.Min()`) of 2 numbers
	fmt.Println("Largest Positive:", math.Max(sixToTheThird, twoToTheFourth))
	fmt.Println("Smallest Positive:", math.Min(sixToTheThird, twoToTheFourth))

	// multiply variables by -1 (now negative values)
	sixToTheThird *= -1
	twoToTheFourth *= -1

	// compare the negative value of the variables and show the largest & smallest numbers
	fmt.Println("Largest Negative:", math.Max(sixToTheThird, twoToTheFourth))
	fmt.Println("Smallest Negative:", math.Min(sixToTheThird, twoToTheFourth))
}

//	% go run main.go
//	Largest Positive: 216
//	Smallest Positive: 16
//	Largest Negative: -16
//	Smallest Negative: -216
