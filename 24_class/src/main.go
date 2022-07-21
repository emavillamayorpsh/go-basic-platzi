package main

import "fmt"

type pc struct {
	ram int
	diks int
	brand string
}

/*
	- String(): this function will be called every time a "pc" struct is called inside a "fmt.print" function
	- String(): is a function that belongs to "struct" so what we are doing here is "overwriting" that function
*/
func (myPc pc) String() string {
	// Sprintf: it creates a string with the result of the values passed by param
	return fmt.Sprintf("I have %d GB RAM , %d GB Disk and is a %s", myPc.ram, myPc.diks, myPc.brand) 
}

func main() {
	myPc := pc{ram: 16, diks: 100, brand: "msi"}

	// fmt.Println(myPc.String()) the String() method is implicit when a struct is called inside a fmt.print function
	// so the code below works exactly as the one defined at the beginning of the comment
	fmt.Println(myPc)

}