package main

import "fmt"

func main() {
	//variables declaration

  helloMessage := "Hello"
  worldMessage := "World"

  // ===========Println===========
  /* it does a print and also adds a "ln" that provokes a "jump" into the next line
  */

  fmt.Println("=====PRINTLN=====")
  fmt.Println(helloMessage, worldMessage)
  fmt.Println(helloMessage, worldMessage)

  // ===========Printf===========
  /*it prints and also adds an extra function to the string that you pass as param 
    - %s: means that there goes a string
    - %d: means that there goes an int
    - %v: means that we don't know the type of the param so we add that
  */
  name := "Platzi"
  courses := 500

  fmt.Println("\n=====PRINTF=====")
  fmt.Printf("%s has more than %d\n", name, courses)
  fmt.Printf("%v has more than %v\n", name, courses)

  //======================
  /* generates a string but it doesn't display it in the console, it returns the value of the string concatenation passed as param
  */
  message := fmt.Sprintf("%s has more than %d", name, courses)
  fmt.Println("\n=====SPRINTF=====")
  fmt.Println("message: ", message)

  // ===========Know Type===========
  /* %t: allows you to know the type of the variable passed as param
  */
  fmt.Println("\n=====KNOW TYPE=====")
  fmt.Printf("hello Message type: %T\n",helloMessage)
  fmt.Printf("courses type: %T\n",courses)
}