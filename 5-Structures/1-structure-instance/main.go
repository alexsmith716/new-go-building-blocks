package main

import (
	"fmt"
)

// 

// create struct type `worker`
// create 3 fields of name/data type pairs (id, name, dept)
type worker struct {
	id   int
	name string
	dept string
}

//	type structName struct {
//		fieldName dataType
//		fieldName dataType
//	}

func main() {

	// create struct instance `writer`
	var writer worker
	// assign data to individual fields
	writer.id = 215
	writer.name = "Elmer"
	writer.dept = "Books"

	// or, using a 'struct literal' (minding order):
	// writer = worker{ 215, "Elmer", "Books" }

	// or, using a 'struct literal' and include field names (any order):
	// writer = worker{ name: "Elmer", id: 215,  dept: "Books" }

	// or, var, instance and 'struct literal' like this:
	// writer := worker{ name: "Elmer", id: 215,  dept: "Books" }

	copywriter := worker{name: "Fudd", dept: "Legal", id: 219}

	fmt.Println(writer)
	fmt.Println(copywriter)

	fmt.Printf("%v works in %v \n", writer.name, writer.dept)
	fmt.Printf("%v works in %v \n", copywriter.name, copywriter.dept)
}

//	% go run main.go
//	{215 Elmer Books}
//	{219 Fudd Legal}
//	Elmer works in Books 
//	Fudd works in Legal
