package main

import "fmt"

func main() {

  // Square Area
  const squareBase = 10
  squareArea := squareBase * squareBase

	fmt.Println("Square Area: ", squareArea)

  x := 10
  y := 50

  // Sum
  result := x + y
  fmt.Println("Sum: ", result)

  // Rest
  result = y - x
  fmt.Println("Rest: ", result)

  // Multiply
  result = x * y
  fmt.Println("Multiply: ", result)

  // Divisor
  result = y / x
  fmt.Println("Divisor: ", result)

  // Module 
  result = y % x
  fmt.Println("Module: ", result)

  // Incremental
  x++
  fmt.Println("Incremental: ", x)

  // Decremental
  x--
  fmt.Println("Decremental: ", x)

}