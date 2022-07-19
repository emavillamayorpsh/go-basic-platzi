package main

import "fmt"

func main() {

  switch module := 5 % 2; module {
  case 0:
    fmt.Println("Is Even")
  default:
    fmt.Println("Is Odd")
  }

  // without condition
  value := 200
  switch  {
  case value > 100: 
    fmt.Println("Is bigger than 100")
  case value < 0:
    fmt.Println("Is less than 0")
  default:
    fmt.Println("No condition")
  }
}