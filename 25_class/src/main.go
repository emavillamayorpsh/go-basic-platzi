package main

import "fmt"

/*INTERFACES:
- The best way to use interfaces is when some struct share the same function
- Interfaces allows us to call shared functions in an easier and reusable way, so that
	we don't need to generate code for each type of struct
*/

type figures2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base float64
	height float64
}

func calculate(f figures2D) {
	fmt.Println("Area: ", f.area())
}

func (c square) area() float64 {
	return c.base * c.base
}

func (r rectangle) area() float64 {
	return r.base * r.height
}

func main() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2 , height: 4}

	calculate(mySquare)
	calculate(myRectangle)


	// Interfaces list
	myInterface := []interface{}{"Hello", 12, 4.90}

	fmt.Println(myInterface...)
}