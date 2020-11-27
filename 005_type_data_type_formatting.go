package main

import (
	f "fmt"
)

//formatting the variable
//ie finding the type of variable declared
//golang being the lang of static type
//where as java script is a dynamic type
var x = 4.0

func main() {
	f.Println(x)
	f.Printf("%T", x)
}
