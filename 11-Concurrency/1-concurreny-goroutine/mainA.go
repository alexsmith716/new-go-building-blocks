package main

import (
	"fmt"
	"time"
)

func learning() {  
	fmt.Println("My first goroutine")
}

func forLoop(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	forLoop("direct")

	go forLoop("goroutine")

	go func(msg string) {
			fmt.Println(msg)
	}("going")

	go learning()

	/* using time sleep so that the main program does not terminate before the execution of goroutine */
	time.Sleep(1 * time.Second)
	//time.Sleep(time.Second)

	fmt.Println("main function")
}


//	Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.

//	To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
//	
//	You can also start a goroutine for an anonymous function call.
//	
//	Our two function calls are running asynchronously in separate goroutines now. Wait for them to finish (for a more robust approach, use a WaitGroup).
//	
//	When we run this program, we see the output of the blocking call first, then the interleaved output of the two goroutines. This interleaving reflects the	 goroutines being run concurrently by the Go runtime.
//	
//	Next we’ll look at a complement to goroutines in concurrent Go programs: channels.