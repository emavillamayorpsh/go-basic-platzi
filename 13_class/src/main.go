package main

import "fmt"

func main() {

  // Conditional For
  fmt.Println("==========CONDITIONAL FOR==========")
  for i := 0; i <= 10; i++ {
    fmt.Print(i," - ")
  }

  // FOR WHILE
  fmt.Println("\n\n==========FOR WHILE==========")
  counter := 0
  for counter <= 10 {
    fmt.Print(counter, " - ")
    counter++
  }

  // FOR FOREVER
  // This for will run forever , that's why is commented
  fmt.Println("\n\n==========FOR FOREVER==========")
  /*   counterForever := 0
  for {
    fmt.Print(counterForever, " - ")
    counterForever++
  } */

  // FOR RANGE
  // is used in for loop to iterate over items of an array, slice, channel or map
  fmt.Println("\n==========FOR RANGE==========")
  evenNumbList:= []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, even := range evenNumbList {
		fmt.Printf("position %d even number: %d \n", i, even)
	}

}