package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
  valueOne := 1
  valueTwo := 2

  if valueOne == 1 {
    fmt.Println("Is 1")
  } else {
    fmt.Println("Is not 1")
  }

  // &&
  if valueOne == 1 && valueTwo == 2 {
    fmt.Println("&&: Is true")
  }

  // ||
  if valueOne == 0 || valueTwo == 2 {
    fmt.Println("||: Is true")
  }

  // Convert text to number
  value, err := strconv.Atoi("53")

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("strconv.Atoi value: ", value)
}