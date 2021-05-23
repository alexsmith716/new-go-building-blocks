package main

import (
	"fmt"
	"math"
)

// https://golang.org/pkg/math/

// 'math' provides functions and constant mathematical values
// constants:				
// math.E, math.Pi, math.Phi, math.Sqrt2, math.SqrtE, math.SqrtPi
// math.SqrtPhi, math.Ln2, math.Log2E, math.Ln10, math.Log10E

// https://golang.org/pkg/fmt/

// %f     * default width, default precision
// %9f    * width 9, default precision
// %.2f   * default width, precision 2
// %9.2f  * width 9, precision 2
// %9.f   * width 9, precision 0

// integers: whole numbers that can be positive, negative, or 0 (..., -1, 0, 1, ...)
// float:    numbers that contain a decimal point, 9.0, -2.25


func main() {

	// 3 floating point (float64) variables
	var radius, area, perimeter float64

	radius = 4
	fmt.Printf("Radius of Circle: %.2f \n", radius)

	area = math.Pi * (radius * radius)
	fmt.Printf("Area of Circle (precision 2): %.2f \n", area)
	fmt.Printf("Area of Circle (precision 4): %.4f \n", area)

	perimeter = 2 * (math.Pi * radius)
	fmt.Printf("Perimeter of Circle (precision 2): %.2f \n", perimeter)
	fmt.Printf("Perimeter of Circle (precision 10): %.10f \n", perimeter)
}

//	% go run main.go
//	Radius of Circle: 4.00 
//	Area of Circle (precision 2): 50.27 
//	Area of Circle (precision 4): 50.2655 
//	Perimeter of Circle (precision 2): 25.13 
//	Perimeter of Circle (precision 10): 25.1327412287
