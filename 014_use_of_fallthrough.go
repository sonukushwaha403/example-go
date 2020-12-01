package main

import (
	"fmt"
)

func main() {

	switch {
	case (2 == 5):
		fmt.Println("wrong")
	case (3 != 3):
		fmt.Println("correct")

	case (4 == 3):
		fmt.Println("wrong")
		fallthrough
	case (3 != 1):
		fmt.Println("wrong")
		fallthrough
	case (1 != 0):
		fmt.Println("wrong")
		
	default:
		fmt.Println("its default")

	}

}
