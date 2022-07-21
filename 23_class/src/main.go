package main

import "fmt"

type pc struct {
	ram int
	disk int
	brand string
}

/*
	- This is a function for the "pc" struct, we let it know it belongs to pc type 
   because of the param in the first parenthesis (before the name of the function)
	- In order to use it we need to have an instance of "pc" and simply access to the function like it is an attribute
*/
func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

/*
 	- The * it means we access to the values of "pc" from the pointer,
		this allows us to pass the value of "pc" by reference so if we change an attribute of pc the change will be also reflected
		outside of the scope of this function (duplicateRAM).
	- If we don't add the "*" to "pc" then pc will be a copy of the struct that is invoking this function so if an attribute is changed
		it won't be reflected outside of the scope of this function because is not pointing or making changes to the same pointer
*/
func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50

	/*to know/access the memory direction where a variable is stored we add the "&" sign at the right of a variable*/
	b := &a // in this case "b" is the pointer of "a" , it means "b" has the memory direction where "a" is stored

	fmt.Println("a memory direction: ", b) // this will print a memory direction value

	/*to know/access the value of a memory direction we add the "*" sign at the right of a variable*/
	fmt.Println("a value getting it from memory direction:", *b) // this will print "50"

	/*in this case we are storing a new value in the memory direction where "a" was stored,
		so now the value of a is "100" instead of 50 */
	*b = 100
	fmt.Println(a)


	myPc := pc{ram: 16, disk: 200, brand: "msi"}

	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)


}