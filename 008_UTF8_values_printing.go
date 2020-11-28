package main

import (
	"fmt"
)

func main() {
	S := 34
	fmt.Printf("%#U\n", S)
	f := "sonu kumar kushwaha"
	for i := 0; i < len(f); i++ {
		fmt.Printf("%#U\n", f[i])
	}
}
