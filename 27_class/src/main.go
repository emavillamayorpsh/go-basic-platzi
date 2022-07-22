package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	// defer: the function passed here it is executed when the rest of the code in the "say" function is executed
	defer wg.Done() // this release the goroutine from this waitGroup
	fmt.Println(text)
}

func main() {
	// sync: allows us to work with goroutines but it is hard to maintain
	var wg sync.WaitGroup // waitGroup: groups go routines and free them little by little

	fmt.Println("Hello")

	// this adds 1 go routine ,in order to wait the execution before the base go routine (main) dies
	wg.Add(1)

	go say("world", &wg) // the reserved word "go" is going to execute the function in a concurrent way

	wg.Wait() // this is the main go routine and we tell them to wait until every go routine is executed

	go func(text string) {
		fmt.Println(text)
	}("Bye")

	time.Sleep(time.Second)
}

/*
	A WaitGroup waits for a collection of goroutines to finish execution.
	For this, bind the WaitGroup.Add() ( wg.add(1) in the class example).
	The integer indicates the number of goroutines to wait to finish executing the main goroutine.

	Every time a goroutine finishes its execution, it calls the Done() method. This causes the WaitGroup counter to decrement.
	When the counter reaches zero, the main routine will continue its execution.

	The wait() function blocks the main routine until all other routines in the group have finished.
*/