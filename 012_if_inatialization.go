package main

import "fmt"

func main() {
	x := 88
	if x := 96; x == 96 {
		fmt.Println("its the inner x", x)
	}
	fmt.Println("its the outer x", x)
}
