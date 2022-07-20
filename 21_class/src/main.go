package main

import "fmt"

/*
	- struct is the "class" for go
	- it has attributes
	- it is basically a definition of fields
*/
type car struct {
	brand string
	year int
}

func main() {
	// instance struct
	/*
		-if you don't add a value to the attributes those will be instanced by the default value of it's type
	*/

	// v1
	myCar := car{brand: "Ford", year: 2022}
	fmt.Println(myCar)

	// v2
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar) // for "year" attribute we didn't specify a value so it will set the default of the type it belongs (int = 0)
}