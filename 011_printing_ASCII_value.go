package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 145; i++ {
		fmt.Printf("%d\t%#U\t%b\n", i, i, i)
	}

}
