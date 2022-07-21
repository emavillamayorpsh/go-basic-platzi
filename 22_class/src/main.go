package main

import (
	"fmt"
	pk "go-basic-platzi/22_class/src/mypackage" // the alias is "pk"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage("Hello ")
}