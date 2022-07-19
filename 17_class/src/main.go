package main

import "fmt"

func main() {
  // Defer
  // before a function finishes it executes the line of code that has "defer" at the beginning
  // it doesn't matter in which position defer is set in the code(beginning or end) it will be executed at the end
  // Use case example: when opening a DB connection you can set in "defer" to close the connection
  /*
  Something important that must be taken into account with the use of the defer
  is that if within the same function you use more than one defer,
  these are treated as a stack or in English "stack" one on top of the other which generates
  that the defers are executed with the LIFO principle (Last In First Out) and you have to be careful in that regard.
  */
  defer fmt.Println("Hello")
  fmt.Println("World")

  // Continue and Break
  for i := 0; i < 10; i++ {
    fmt.Println(i)

    // Continue
    /*
    The continue statement is used when you want to skip the remaining part of the loop,
    return to the top of the loop, and continue with a new iteration.
    */
    if(i == 2) {
      fmt.Println("Is 2")
      continue
    }

    // Break
    /*
    The break statement is used when you want to get out of the loop and no iterate over it anymore
    */
    if i == 8 {
      fmt.Println("Break")
      break
    }
  }
}