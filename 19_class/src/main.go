package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	parsedText := strings.ToLower(text)
	var textReverse string

	for i:= len(parsedText) -1; i >= 0; i-- {
		// we wrap the text with "string" because if we exceed the length of the "text" it will return the ascii code
		textReverse += string(parsedText[i])
	}

	if parsedText == textReverse {
		fmt.Println("Is palindromo")
	} else {
		fmt.Println("Is not palindromo")
	}
}

func main() {
	slice :=[]string {"hello", "whats", "up"}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo("Neuquen")
}