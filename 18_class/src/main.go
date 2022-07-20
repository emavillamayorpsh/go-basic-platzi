package main

import "fmt"

func main() {
	// Array
	/*
		- format: var/const name [size] type
		- The spaces of the array have default values in case they aren't set , similar to a common const/variable that doesn't
			have a value specified
		- You must specify the size of the array at the creation
		- immutable: you can't add more element to the array once the size is defined
		- Functions:
			* len(array): this function let us know how many elements an array have
			* cap(array): this function let us know the maximum capacity of the array
	*/
	var array [4]int
	array[0] = 1
	array[1] = 2

	fmt.Println("=============ARRAY=============")
	fmt.Println("array: ", array)
	fmt.Println("len(array): ", len(array))
	fmt.Println("cap(array): ", cap(array))

	// Slice
	/*
		- It doesn't have a size specified
		- mutable: you can add elements to the slice
	*/

	slice := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println("\n=============SLICE=============")
	fmt.Println("slice: ", slice)
	fmt.Println("len(slice): ", len(slice))
	fmt.Println("cap(slice): ", cap(slice))


	// Methods in slice
	fmt.Println("\n=============METHODS IN SLICE=============")
	fmt.Println(slice[0])
	fmt.Println(slice[:3]) // it will print "until the third position" that means from 0 to 3 position
	fmt.Println(slice[2:4]) // it will print the elements between 2 and 4 position
	fmt.Println(slice[4:]) // it will print "from the fourth position onwards" that means from 4 position to the last position

	// Append
	fmt.Println("\n=============APPEND METHOD=============")
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append new list
	fmt.Println("\n=============APPEND NEW LIST=============")
	newSlice := []int{8, 9, 10}

	/*the 3 dots (...) at the right of "newSlice" means that it will "decompress" the array and get the elements inside of it 
	and put those inside of the "slice" slice*/
	slice = append(slice, newSlice...)
	fmt.Println(slice)


}