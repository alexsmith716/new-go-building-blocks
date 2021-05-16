package main

import (
	"fmt"
)

func main() {

	var a, b int

	a, b = 8, 4
	fmt.Println("Assigned Values:\ta =", a, "\tb =", b)

	a += b // a = 8 + 4.
	fmt.Println("Add and Assign:\t\ta =", a)
	a -= b // a = 12 - 4.
	fmt.Println("Subtract and Assign:\ta =", a)
	a *= b // a = 8 * 4.
	fmt.Println("Multiply and Assign:\ta =", a)
	a /= b // a = 32 / 4.
	fmt.Println("Divide and Assign:\ta =", a)
	a %= b // a = 8 % 4.
	fmt.Println("Remainder Assigned:\ta =", a)
}

//	nothing new here
//	operator
//	=					a = b      a = b
//	+=				a += b     a = (a+b)
//	-=				a -= b     a = (a-b)
//	*=				a *= b     a = (a*b)
//	/=				a /= b     a = (a/b)
//	%=				a %= b     a = (a%b)

//	% go run main.go
//	Assigned Values:	a = 8 	b = 4
//	Add and Assign:		a = 12
//	Subtract and Assign:	a = 8
//	Multiply and Assign:	a = 32
//	Divide and Assign:	a = 8
//	Remainder Assigned:	a = 0
