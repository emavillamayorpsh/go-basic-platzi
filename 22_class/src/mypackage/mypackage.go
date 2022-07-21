package mypackage

import "fmt"

/*
	PACKAGE:
	 - All identifiers defined within a package are visible throughout that package.
	 - When importing a package you can access only its exported identifiers.
	 - An identifier is exported if it begins with a capital letter.
	 - the attributes that starts with a smallCase letter are private so you can access to it
	 	 you get an instance of the struct
*/

// CarPublic Car with public access
type CarPublic struct {
	Brand string
	Year int
}

// the name of the struct starts in smallCase because is private
type carPrivate struct {
	brand string
	year int
}

// PrintMessage() prints a message
func PrintMessage(text string) {
	fmt.Println(text)
}

// this function can't be exported because it starts with a smallCase and it's private
func printMessageOther(text string) {
	fmt.Println(text)
}