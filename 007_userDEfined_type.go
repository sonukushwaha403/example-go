package main

import (
	"fmt"
)
var x int
type y int
var z y
func main() {
	x=34
	z=45
	//in golang ever thing is explicit type 
	//there is nothing called implicit type
	x=int(z)
	fmt.Print(" z ,its is of type ",z)
	fmt.Printf(" %T\n",z)
	fmt.Println("Hello, playground",    x)
}
