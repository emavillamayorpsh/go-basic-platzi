// first step : declare a package, it means the name of the folder where it is stored
// in this case since this is the main file , the name of the package is "main"
package main

//fmt: this package is in charge of console print
import (
	"fmt"
	"reflect"
)

//second step: declare main function that will execute the code
func main(){
  /*===============Declaration of constants===============*/
  // const name type = value
  const pi float64 = 3.14
  // if don't specify the type then it will assume once by default , in this case it assumes "float64"
  const pi2 = 3.14

  fmt.Println("===CONSTANTS===")
  fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)
  fmt.Println("Pi2 type:", reflect.TypeOf(pi2))

  /*===============Declaration of variables===============*/
  /*GO WON'T COMPILE IF VARIABLES AREN'T USE*/
  //1st way: if we haven't use a variable before we can define that by adding ":" before the "=" symbol
  // so it will CREATE the variable and ASSIGN the value like the following example
  base := 12

  //2nd way: in this other way we declare the variable (var) and assign a value to it but we need to define the TYPE (int) 
  var height int = 14

  //3rd way: in this case we create the declare the variable but we don't assign any value to it
  var area int

  fmt.Println("\n===VARIABLES===")
  fmt.Println(base, height, area)

  /*===============Zero values===============*/
  /*It is a value that it will have a variable if we don't add a value , go assign a default value to variables */
  var a int
  var b float64
  var c string
  var d bool


  fmt.Println("\n===ZERO VALUES===")
  fmt.Println("int: ", a) // a: 0
  fmt.Println("float64: ", b) // b: 0
  fmt.Println("c: ", c) // c: (empty string)
  fmt.Println("d: ", d) // false

  // square area
  const squareBase = 10
  squareArea := squareBase * squareBase

  fmt.Println("\nSquare Area: ", squareArea)

}

/* Two different ways to run the code:
  1- go build main.go  -> this way a "main" file will be created and it can be executed with "./main" and it will print "Hello world"
  2- go run main.go    -> this way no file is created and it runs the code immediately but it not as efficient as "go build"
*/
