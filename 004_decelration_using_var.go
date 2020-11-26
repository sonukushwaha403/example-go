package main

import (
	s "fmt"
)
var y = "sonu"
func main() {
	s.Println("s is alias for fmt ")
	//identifier deceleration with var keyword can be done outside function body
	s.Println(y)
	//short hand decelration is only possible inside function body 
	x:=34
	s.Println(x)
}
