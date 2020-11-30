package main

import (
	"fmt"
)

func main() {

	for i := 22; i <= 124; i++ {
		fmt.Printf("%d\t%#U\t%b\n", i, i, i)
	}

}
