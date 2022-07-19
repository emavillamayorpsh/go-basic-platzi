package main

import "fmt"

func normalFunction(message string) {
  fmt.Println(message)
}

// v1 declare params types
// func tripleArgument(a int, b int, c string) {

// v2 declare params types
/* if a param is the same type of the other , there is no need to write again the type , just separate both params with a "," and
   add the type at the end
*/
func tripleArgument(a, b int, c string) {
  fmt.Println(a, b, c)
}

// return one value
func returnValue(a int) int {
  return a * 2
}

// return one or more values
func doubleReturnValue(a int) (c, d int) {
  return a, a * 2
}

func main() {
  normalFunction("Hello world")

  tripleArgument(1, 2, "Hello")

  value := returnValue(2)
  fmt.Println("Value: ", value)

  // get two values of a function that returns two values
  valueOne, valueTwo := doubleReturnValue(2)
  fmt.Printf("valueOne: %d - valueTwo: %d\n", valueOne, valueTwo)

  // get one value of a function that returns two values
  // for this case we add "_" to avoid getting another value
  justValueOne, _ := doubleReturnValue(2)
  fmt.Println("Just value one: ", justValueOne)
}