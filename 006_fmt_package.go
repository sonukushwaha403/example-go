package main

import (
	f "fmt"
)

var x int= 34
var y = " its identifier name is y"

func main() {
	//x := 34
	f.Printf("%b\n", x)
	f.Println(x)
	f.Println("sonu kumar kushwaha" , x)
	z := f.Sprint(x)
	f.Println(z)
}
