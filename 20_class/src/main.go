package main

import "fmt"

func main() {

	// DICTIONARY
	/*
		- format: nameOfTheVariable := make(map[keyType] valueType)
		- Usually the "keyType" is string
		- make() allows us to create dictionaries
	*/
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Loop map
	fmt.Println("============LOOP MAP============")
	/*
		- take into account that when looping over a dictionary it won't always do it in the same way the elements were created
	*/
	for i,v := range m {
		fmt.Println(i,v)
	}

	// FIND A VALUE
	fmt.Println("============FIND A VALUE============")

	/* the "ok" value indicates if the "key" that we are looking in the dictionary exists
		because if the key doesn't exist "value" has the default value of the type of the value 
		in this case since the values of the keys are "int" it will store a "0"
	*/
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}